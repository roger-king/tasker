import React from 'react';
import { Box, Select } from 'grommet';
import styled from 'styled-components';

import CategoryList from '../components/categoryList';
import TaskList from '../components/taskList';

import { listTasks } from '../data/tasker';

interface HomePageProps {
    className?: string;
}

interface HomePageState {
    currentCategory: string | null;
    categories: string[];
    tasks: any[];
}

class HomePage extends React.PureComponent<HomePageProps, HomePageState> {
    constructor(props: HomePageProps) {
        super(props);

        this.state = {
            currentCategory: null,
            categories: [],
            tasks: [],
        };
    }

    async componentDidMount(): Promise<void> {
        const tasks = await listTasks();
        const categories = tasks.data.map((t: any) => t.executor);
        this.setState({ categories, tasks: tasks.data });
    }

    setCurrentCategory = (c: string): void => {
        this.setState({ currentCategory: c });
    };

    render(): JSX.Element {
        const { className } = this.props;
        const { currentCategory, categories, tasks } = this.state;
        return (
            <Box className={className} direction="row" gap="large" justify="between" fill pad="xlarge">
                <CategoryList
                    categories={categories}
                    current={currentCategory}
                    selectCategory={this.setCurrentCategory}
                />
                <Box align="center" pad={{ left: '70px' }} width="100%">
                    {tasks.length === 0 ? 'loading...' : <TaskList header="main" tasks={tasks} />}
                </Box>
                <Box direction="row" align="start" justify="center" gap="small">
                    <Select options={['Last 7 days']} size="small" value="Last 7 days" />
                    <Select options={['All', 'main']} size="small" value="All" />
                </Box>
            </Box>
        );
    }
}

export default styled(HomePage)``;
