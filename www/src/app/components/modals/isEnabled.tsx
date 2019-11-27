import React from 'react';
import { Box, Button, Text } from 'grommet';
import Modal from './index';
import { capitalize } from '../../utils/case';

import { disableTask } from '../../data/tasker';

interface IsEnabledModalProps {
    showModal: any;
    id: string;
    name: string;
    enabled: boolean;
}

const IsEnabledModal: React.FC<IsEnabledModalProps> = (props: IsEnabledModalProps): JSX.Element => {
    const { showModal, name, id, enabled } = props;
    const header = enabled ? 'disable' : 'enable';

    return (
        <Modal setShowModal={showModal} width="medium" header={capitalize(header)} onClickOutside={false} onEsc>
            <Box>
                <Box margin="small" direction="row" gap="small">
                    <Text>
                        Are you sure you would like to {header} your task?
                        <Box pad="xsmall" background="light-4">
                            <pre>{name}</pre>
                        </Box>
                    </Text>
                </Box>
                <Box margin="small" direction="row" gap="small">
                    <Button
                        color="light-4"
                        label="Cancel"
                        onClick={(): void => showModal(false)}
                        style={{ borderRadius: '4px' }}
                    />
                    <Button
                        primary
                        color="brand"
                        label={capitalize(header)}
                        onClick={async (): Promise<void> => {
                            if (enabled) {
                                await disableTask(id);
                            } else {
                                // await disableTask(id);
                            }

                            showModal(false);
                        }}
                        style={{ borderRadius: '4px' }}
                    />
                </Box>
            </Box>
        </Modal>
    );
};

export default IsEnabledModal;
