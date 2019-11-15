import React, { useState } from 'react';
import { Layer, Box, Button, Heading, Text, TextInput } from 'grommet';
import { Alert, Close } from 'grommet-icons';

import { deleteTask } from '../../data/tasker';

interface DeleteTaskModalProps {
    showModal: any;
    taskId: string;
    taskName: string;
}

const DeleteTaskModal: React.FC<DeleteTaskModalProps> = (props: DeleteTaskModalProps): JSX.Element => {
    const { showModal, taskName, taskId } = props;
    const [confirmedName, setConfirmedName] = useState('');
    const [isDisabledDestroyBtn, toggleCanDestroy] = useState(true);

    return (
        <Layer modal onClickOutside={(): void => showModal(false)} onEsc={(): void => showModal(false)}>
            <Box width="medium" pad="medium">
                <Box direction="row">
                    <Button icon={<Close size="medium" />} onClick={(): void => showModal()} />
                    <Heading level="4">Delete Task</Heading>
                </Box>
                <Box margin="small" direction="row" gap="small">
                    <Text>
                        <Alert color="warning" />
                        <b>WARNING:</b> Deleting {taskName} will remove any pending actions that may not have occured
                        yet.
                    </Text>
                </Box>
                <Box margin="small">
                    <Text>Confirm you want to permanently delete {taskName} by entering its name below.</Text>
                    <Box
                        margin="small"
                        background="light-2"
                        height="50px"
                        border="all"
                        align="start"
                        justify="center"
                        pad={{ left: '12px' }}
                        style={{ borderRadius: '4px' }}
                    >
                        {taskName}
                    </Box>
                    <Box margin="small" height="45px">
                        <TextInput
                            placeholder="Enter the name of the task"
                            value={confirmedName}
                            onChange={(e: any) => {
                                setConfirmedName(e.target.value);
                                if (e.target.value === taskName) {
                                    toggleCanDestroy(false);
                                } else {
                                    toggleCanDestroy(true);
                                }
                            }}
                        />
                    </Box>
                    <Box direction="row-responsive" gap="small" margin="small">
                        <Button
                            primary
                            label="Cancel"
                            style={{ borderRadius: '4px' }}
                            onClick={() => showModal(false)}
                            color="light-4"
                        />
                        <Button
                            primary
                            color="warning"
                            label="Destroy"
                            disabled={isDisabledDestroyBtn}
                            style={{ borderRadius: '4px' }}
                            onClick={async () => {
                                await deleteTask(taskId);
                                showModal(false);
                            }}
                        />
                    </Box>
                </Box>
            </Box>
        </Layer>
    );
};

export default DeleteTaskModal;
