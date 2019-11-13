import React from 'react';
import styled from 'styled-components';
import { Box, Heading } from 'grommet';
import Task from './task';

interface TaskListProps {
    className?: string;
    header: string;
    tasks: any[];
}

const TaskList: React.FC<TaskListProps> = (props: TaskListProps): JSX.Element => {
    const { className, header, tasks } = props;

    return (
        <Box className={className} fill>
            <Box border="bottom" width="100%">
                <Heading level="3" margin="xsmall">
                    {header}.
                </Heading>
            </Box>
            <Box margin="small" gap="small">
                {tasks.map((t: any) => (
                    <Task
                        key={t.name}
                        name={t.name}
                        description={t.description}
                        enabled={t.enabled}
                        complete={t.complete}
                        runTime={t.runTime}
                    />
                ))}
            </Box>
        </Box>
    );
};

export default styled(TaskList)``;
