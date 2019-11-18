import React from 'react';
import { Box, Button, Heading, Layer, Text } from 'grommet';
import { Close } from 'grommet-icons';
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
        <Layer modal onClickOutside={(): void => showModal(false)} onEsc={(): void => showModal(false)}>
            <Box width="medium" pad="medium">
                <Box direction="row">
                    <Button icon={<Close size="medium" />} onClick={(): void => showModal(false)} />
                    <Heading level="4">{capitalize(header)} Task</Heading>
                </Box>
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
        </Layer>
    );
};

export default IsEnabledModal;
