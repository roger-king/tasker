import React from 'react';
import { Box, Select } from 'grommet';
import styled from 'styled-components';

import CategoryList from '../components/categoryList';
import TaskList from '../components/taskList';

interface HomePageProps {
    className?: string;
}

const HomePage: React.FC<HomePageProps> = (props: HomePageProps): JSX.Element => {
    const { className } = props;
    return (
        <Box className={className} direction="row" gap="large" justify="between" fill pad="xlarge">
            <CategoryList categories={['main', 'create', 'delete']} />
            <Box align="center" pad={{ left: '70px' }} width="100%">
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
            <Box direction="row" align="start" justify="center" gap="small">
                <Select options={['Last 7 days']} size="small" value="Last 7 days" />
                <Select options={['All', 'main']} size="small" value="All" />
            </Box>
        </Box>
    );
};

export default styled(HomePage)``;
