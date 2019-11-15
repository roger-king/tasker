import React from 'react';
import { Box, Select } from 'grommet';
import styled from 'styled-components';

import CategoryList from '../components/categoryList';
import TaskList from '../components/taskList';
import CreateTaskModal from '../components/modals/createTask';

import { listTasks } from '../data/tasker';

interface HomePageProps {
    className?: string;
}

interface HomePageState {
    currentCategory: string | null;
    categories: string[];
    tasks: Task[];
    filteredTasks: Task[];
    showModal: boolean;
}

class HomePage extends React.PureComponent<HomePageProps, HomePageState> {
    constructor(props: HomePageProps) {
        super(props);

        this.state = {
            currentCategory: null,
            categories: [],
            tasks: [],
            filteredTasks: [],
            showModal: false,
        };
    }

    async componentDidMount(): Promise<void> {
        const tasks = await listTasks();
        const categories: string[] = [];

        tasks.data.map((t: Task): void => {
            if (categories.indexOf(t.executor) === -1) {
                categories.push(t.executor);
            }
            // eslint-disable-next-line
            return;
        });

        this.setState({ categories, tasks: tasks.data });
    }

    setCurrentCategory = (c: string | null): void => {
        const { tasks } = this.state;
        let currTasks: any[] = [];

        if (c === null) {
            currTasks = [];
        } else {
            currTasks = tasks.filter((t: Task) => t.executor === c);
        }

        this.setState({ currentCategory: c, filteredTasks: currTasks });
    };

    setShowModal = (): void => {
        const { showModal } = this.state;
        this.setState({ showModal: !showModal });
    };

    render(): JSX.Element {
        const { className } = this.props;
        const { currentCategory, categories, tasks, filteredTasks, showModal } = this.state;
        return (
            <Box className={className} direction="row" gap="large" justify="between" fill pad="xlarge">
                <CategoryList
                    categories={categories}
                    current={currentCategory}
                    selectCategory={this.setCurrentCategory}
                    openModal={this.setShowModal}
                />
                <Box align="center" pad={{ left: '70px' }} width="100%">
                    {tasks.length === 0 ? (
                        'loading...'
                    ) : (
                        <TaskList
                            tasks={filteredTasks.length === 0 ? tasks : filteredTasks}
                            categories={categories}
                            currentCategory={currentCategory}
                        />
                    )}
                </Box>
                <Box direction="row" align="start" justify="center" gap="small">
                    <Select options={['Last 7 days']} size="small" value="Last 7 days" />
                    <Select
                        options={['All', ...categories]}
                        size="small"
                        value={currentCategory || 'All'}
                        onChange={({ option }): void => {
                            if (option === 'All') {
                                this.setCurrentCategory(null);
                            } else {
                                this.setCurrentCategory(option);
                            }
                        }}
                    />
                </Box>
                {showModal && <CreateTaskModal showModal={this.setShowModal} />}
            </Box>
        );
    }
}

export default styled(HomePage)``;
