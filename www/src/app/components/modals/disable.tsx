import React from 'react';
import { Box, Button, Heading, Layer } from 'grommet';
import { Close } from 'grommet-icons';

interface DisableModalProps {
    showModal: any;
    id: string;
    name: string;
}

const DisableModal: React.FC<DisableModalProps> = (props: DisableModalProps): JSX.Element => {
    const { showModal } = props;

    return (
        <Layer modal onClickOutside={(): void => showModal(false)} onEsc={(): void => showModal(false)}>
            <Box width="medium" pad="medium">
                <Box direction="row">
                    <Button icon={<Close size="medium" />} onClick={(): void => showModal(false)} />
                    <Heading level="4">Disable Task</Heading>
                </Box>
            </Box>
        </Layer>
    );
};

export default DisableModal;
