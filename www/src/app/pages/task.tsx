import React, { useState, useEffect } from 'react';
import { Box, Heading } from 'grommet';

import { RouteComponentProps } from 'react-router';
import { findTask } from '../data/tasker';

interface TaskPageProps extends RouteComponentProps {
    className?: string;
}

const TaskPage: React.FC<TaskPageProps> = (props: TaskPageProps): JSX.Element => {
    const { match } = props;
    const { params }: any = match;
    const [task, setTask] = useState<Task | null>(null);

    useEffect(() => {
        const fetchTask = async (): Promise<void> => {
            const { id } = params;
            const foundTask = await findTask(id);

            setTask(foundTask.data);
        };

        fetchTask();
    }, [params]);

    if (task) {
        return (
            <Box align="center" fill>
                <Heading>{task.name}</Heading>
            </Box>
        );
    }

    return <Box>loading...</Box>;
};

export default TaskPage;
