import React from 'react';
import { Box, Button, Heading, Layer, Text } from 'grommet';
import { Close } from 'grommet-icons';

interface DisableModalProps {
    showModal: any;
    id: string;
    name: string;
}

const DisableModal: React.FC<DisableModalProps> = (props: DisableModalProps): JSX.Element => {
    const { showModal, name } = props;

    return (
        <Layer modal onClickOutside={(): void => showModal(false)} onEsc={(): void => showModal(false)}>
            <Box width="medium" pad="medium">
                <Box direction="row">
                    <Button icon={<Close size="medium" />} onClick={(): void => showModal(false)} />
                    <Heading level="4">Disable Task</Heading>
                </Box>
                <Box margin="small" direction="row" gap="small">
                    <Text>
                        Are you sure you would like to disable your task?
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
                        label="Disable"
                        onClick={(): void => console.log('disable')}
                        style={{ borderRadius: '4px' }}
                    />
                </Box>
            </Box>
        </Layer>
    );
};

export default DisableModal;
