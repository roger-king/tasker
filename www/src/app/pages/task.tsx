import React, { useState, useEffect } from 'react';
import { Box, Heading, Button, Text } from 'grommet';
import { Edit } from 'grommet-icons';
import { DateTime } from 'luxon';
import Parser from 'cron-parser';
import { RouteComponentProps } from 'react-router';
import PrettyJSON from '../components/prettyJSON';
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
        const formattedCreatedAt = DateTime.fromISO(String(task.createdAt)).toLocaleString(DateTime.DATETIME_MED);
        const schedule = Parser.parseExpression(task.schedule)
            .next()
            .toString();

        return (
            <Box align="start" gap="medium" fill>
                <Box>
                    <Box align="center" alignSelf="start" direction="row" gap="small" width="100%">
                        <Box gap="small">
                            <Heading level="2" margin="none">
                                {task.name}
                            </Heading>
                            <Text size="16px">
                                Created at <b>{formattedCreatedAt}</b>
                            </Text>
                        </Box>
                        <Box direction="row" gap="small" align="center" alignSelf="start">
                            <Button icon={<Edit size="small" />} label="Edit" style={{ borderRadius: '8px' }} />
                        </Box>
                    </Box>
                    <Box margin="medium">
                        <Text>
                            <i>{task.description}</i>
                        </Text>
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
                    <Button primary color="brand" label="Disable" style={{ borderRadius: '8px' }} />
                </Box>
            </Box>
        );
    }

    return <Box>loading...</Box>;
};

export default TaskPage;
