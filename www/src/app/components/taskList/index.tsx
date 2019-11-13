import React from 'react';
import styled from 'styled-components';
import { Box, Heading } from 'grommet';
import Task from './task';

interface TaskListProps {
    className?: string;
    tasks: Task[];
    categories: string[];
}

interface TaskListData {
    [key: string]: Task[];
}

const TaskList: React.FC<TaskListProps> = (props: TaskListProps): JSX.Element => {
    const { className, tasks, categories } = props;

    return (
        <Box className={className} gap="small" fill>
            {categories.map(
                (c: string): JSX.Element => {
                    return (
                        <Box key={c} gap="small">
                            <Box border="bottom" width="100%">
                                <Heading level="3" margin="xsmall">
                                    {c}.
                                </Heading>
                            </Box>
                            {tasks.map((t: Task): JSX.Element | null => {
                                if (t.executor === c) {
                                    return (
                                        <Task
                                            key={t.taskId}
                                            name={t.name}
                                            description={t.description}
                                            enabled={t.enabled}
                                            complete={t.complete}
                                            runTime={t.schedule}
                                        />
                                    );
                                }
                                return null;
                            })}
                        </Box>
                    );
                },
            )}
        </Box>
    );
};

export default styled(TaskList)``;
