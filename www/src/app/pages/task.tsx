import React, { useState, useEffect } from 'react';
import { Box, Heading, Button, Text } from 'grommet';
import { Edit } from 'grommet-icons';
import { DateTime } from 'luxon';
import Parser from 'cron-parser';
import { RouteComponentProps } from 'react-router';

import PrettyJSON from '../components/prettyJSON';
import TaskTypeIcon from '../components/taskTypeIcon';
import { findTask } from '../data/tasker';
import DisableModal from '../components/modals/disable';

interface TaskPageProps extends RouteComponentProps {
    className?: string;
}

const TaskPage: React.FC<TaskPageProps> = (props: TaskPageProps): JSX.Element => {
    const { match } = props;
    const { params }: any = match;
    const [task, setTask] = useState<Task | null>(null);
    const [showModal, setShowModal] = useState<boolean>(false);

    useEffect(() => {
        const fetchTask = async (): Promise<void> => {
            const { id } = params;
            const foundTask = await findTask(id);

            setTask(foundTask.data);
        };

        fetchTask();
    }, [params]);

    if (task) {
        const formattedCreatedAt = DateTime.fromISO(String(task.createdAt)).toLocaleString(DateTime.DATETIME_MED);
        const schedule = Parser.parseExpression(task.schedule)
            .next()
            .toString();

        return (
            <Box gap="medium" fill>
                <Box margin={{ top: '80px', left: '100px', right: '100px' }} gap="medium">
                    <Box align="center" direction="row" gap="small" width="100%" justify="between">
                        <Box gap="small">
                            <Box direction="row" align="center" gap="small">
                                <TaskTypeIcon repeat={task.isRepeatable} />
                                <Heading level="2" margin="none" color="accent-1">
                                    {task.name}
                                </Heading>
                            </Box>
                            <Text>{task.description}</Text>
                            <Text size="16px">
                                <i>
                                    Created at <b>{formattedCreatedAt}</b>
                                </i>
                            </Text>
                        </Box>
                        <Button
                            icon={<Edit size="small" />}
                            label="Edit"
                            style={{ borderRadius: '8px' }}
                            alignSelf="start"
                        />
                    </Box>
                    <Box align="center" width="100%">
                        <PrettyJSON data={task.args} />
                    </Box>
                </Box>
                <Box
                    flex={false}
                    height="100px"
                    width="100%"
                    align="center"
                    justify="center"
                    direction="row"
                    background="accent-1"
                    style={{ position: 'absolute', bottom: 0 }}
                    gap="medium"
                >
                    <Text>{schedule}</Text>
                    <Button
                        primary
                        color="brand"
                        label="Disable"
                        style={{ borderRadius: '8px' }}
                        onClick={(): void => setShowModal(true)}
                    />
                </Box>
                {showModal && <DisableModal showModal={setShowModal} id={task.taskId} name={task.name} />}
            </Box>
        );
    }

    return <Box>loading...</Box>;
};

export default TaskPage;
