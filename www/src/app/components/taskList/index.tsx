import React from 'react';
import styled from 'styled-components';
import { Box, Heading } from 'grommet';
import Task from './task';

interface TaskListProps {
    className?: string;
    tasks: Task[];
}

const TaskList: React.FC<TaskListProps> = (props: TaskListProps): JSX.Element => {
    const { className, tasks } = props;
    const headers: string[] = [];

    return (
        <Box className={className} fill>
            {tasks.map((t: Task) => (
                <Box key={t.name} margin="small" gap="small">
                    {headers.indexOf(t.executor) === -1 ? (
                        <Box border="bottom" width="100%">
                            <Heading level="3" margin="xsmall">
                                {t.executor}.
                            </Heading>
                        </Box>
                    ) : (
                        ''
                    )}
                    <Task
                        name={t.name}
                        description={t.description}
                        enabled={t.enabled}
                        complete={t.complete}
                        runTime={t.schedule}
                    />
                </Box>
            ))}
        </Box>
    );
};

export default styled(TaskList)``;
