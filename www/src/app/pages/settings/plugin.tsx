import React, { useState, useEffect } from 'react';
import { Box, Heading, Text, CheckBox, Table, TableRow, TableCell, TableBody, Button } from 'grommet';
import { StatusInfo } from 'grommet-icons';
import { listSettings, toggleActivePluginSetting } from '../../data/settings';
import Modal from '../../components/modals';

interface PluginInfoModalProps {
    setShowModal: any;
}

const PluginInfoModal: React.FC<PluginInfoModalProps> = (props: PluginInfoModalProps) => {
    const { setShowModal } = props;

    return (
        <Modal setShowModal={setShowModal} header="Info" width="medium" onClickOutside onEsc>
            <Box>testing</Box>
        </Modal>
    );
};

const PluginSettingPage: React.FC = () => {
    const [loading, setLoading] = useState<boolean>(true);
    const [showInfoModal, setShowInfoModal] = useState<boolean>(false);
    const [settings, setSettings] = useState<Setting[] | null>(null);
    const [error, setError] = useState<any>();

    const getSettings = () => {
        listSettings()
            .then(({ data }: { data: Setting[] | null }) => {
                setLoading(false);
                setSettings(data);
            })
            .catch((err: any) => {
                setError(err);
            });
    };

    useEffect(() => {
        getSettings();
    }, []);

    if (error) {
        return <Box>{error}</Box>;
    }

    if (loading) {
        return <Box>loading...</Box>;
    }

    if (settings && settings.length > 0) {
        return (
            <Box>
                <Heading level="4">Repositories:</Heading>
                <Text size="14px">a place for tasker to find your scripts/plugins</Text>
                <Box margin={{ top: '20px' }} gap="medium">
                    <Box pad="15px" background="light-3" direction="row" align="center" justify="between">
                        <Button label="add repository" primary style={{ borderRadius: '7px' }} />
                        <Text>{settings.length} Repositories</Text>
                    </Box>
                    <Table>
                        <TableBody>
                            {settings.map((s: Setting) => (
                                <TableRow key={s.repo_name}>
                                    <TableCell>
                                        <CheckBox
                                            toggle
                                            checked={s.active}
                                            onChange={async (): Promise<void> => {
                                                const done = await toggleActivePluginSetting({
                                                    repo_name: s.repo_name,
                                                    active: !s.active,
                                                });

                                                if (done) {
                                                    getSettings();
                                                }
                                            }}
                                        />
                                    </TableCell>
                                    <TableCell>{s.repo_name}</TableCell>
                                    <TableCell>
                                        <Button icon={<StatusInfo />} onClick={(): void => setShowInfoModal(true)} />
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </Box>
                {/* Modals */}
                {showInfoModal && <PluginInfoModal setShowModal={setShowInfoModal} />}
            </Box>
        );
    }

    return <Box>configure your plugin</Box>;
};

export default PluginSettingPage;
