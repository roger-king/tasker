import React from 'react';
import { Box } from 'grommet';
import styled from 'styled-components';

import CategoryList from '../components/categoryList';
import TaskList from '../components/taskList';

interface HomePageProps {
    className?: string;
}

const HomePage: React.FC<HomePageProps> = (props: HomePageProps): JSX.Element => {
    const { className } = props;
    return (
        <Box className={className} direction="row">
            <CategoryList categories={['main', 'create', 'delete']} />
            <Box align="start" width="80%" pad={{ left: '200px', right: '200px' }}>
                <TaskList
                    header="main"
                    tasks={[
                        {
                            name: 'test',
                            description: 'a description of a test',
                            set: true,
                            complete: false,
                            runTime: 'May 29, 2020',
                        },
                        {
                            name: 'create application user',
                            description: 'creating an application user for RDS instances',
                            set: true,
                            complete: false,
                            runTime: 'May 29, 2020',
                        },
                    ]}
                />
            </Box>
        </Box>
    );
};

export default styled(HomePage)``;
