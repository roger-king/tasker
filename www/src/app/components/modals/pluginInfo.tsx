import React from 'react';
import { Box, Anchor, Text, Heading } from 'grommet';

import Modal, { BaseModalProps } from './index';

interface PluginInfoModalProps extends BaseModalProps {
    setSetting: any;
    setting: PluginSetting;
}

const PluginInfoModal: React.FC<PluginInfoModalProps> = (props: PluginInfoModalProps) => {
    const { setShowModal, setSetting, setting } = props;

    return (
        <Modal
            // eslint-disable-next-line
            setShowModal={(toggle: boolean = false): void => {
                setSetting(null);
                setShowModal(toggle);
            }}
            header="Info"
            width="medium"
            onClickOutside
            onEsc
        >
            <Box gap="small">
                <Box border="bottom">
                    <Heading level="4" margin="xsmall">
                        {setting.repo_name}
                    </Heading>
                </Box>
                <Text wordBreak="keep-all">
                    <i>{setting.description ? setting.description : 'no description provided'}</i>
                </Text>
                <Box margin={{ top: '30px' }}>
                    <Anchor href={`https://github.com/${setting.repo_name}`} target="_blank">
                        View Github
                    </Anchor>
                </Box>
            </Box>
        </Modal>
    );
};

export default PluginInfoModal;
