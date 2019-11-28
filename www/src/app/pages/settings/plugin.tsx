import React, { useState, useEffect } from 'react';
import { Box, Heading, Text, CheckBox, Table, TableRow, TableCell, TableBody, Button } from 'grommet';
import { StatusInfo } from 'grommet-icons';
import { listSettings, toggleActivePluginSetting } from '../../data/settings';
import AddPluginModal from '../../components/modals/createPlugin';
import PluginInfoModal from '../../components/modals/pluginInfo';

const PluginSettingPage: React.FC = () => {
    const [loading, setLoading] = useState<boolean>(true);
    const [showInfoModal, setShowInfoModal] = useState<boolean>(false);
    const [showAddModal, setShowAddModal] = useState<boolean>(false);
    const [settings, setSettings] = useState<Setting[] | null>(null);
    const [selectedSetting, setSelectedSetting] = useState<Setting | null>(null);
    const [error, setError] = useState<any>();

    const getSettings = (): void => {
        listSettings()
            .then(({ data }: { data: Setting[] | null }) => {
                setLoading(false);
                setSettings(data);
            })
            .catch((err: any) => {
                console.error(err);
                setError('An error occured fetching for settings');
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
                        <Button label="add repository" primary onClick={(): void => setShowAddModal(true)} />
                        <Text>{settings.length} Repositories</Text>
                    </Box>
                    <Table>
                        <TableBody>
                            {settings.map((s: Setting) => (
                                <TableRow key={s.repo_name}>
                                    <TableCell border="bottom">
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
                                    <TableCell border="bottom"> {s.repo_name}</TableCell>
                                    <TableCell border="bottom">
                                        <Button
                                            icon={<StatusInfo />}
                                            onClick={(): void => {
                                                setSelectedSetting(s);
                                                setShowInfoModal(true);
                                            }}
                                        />
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </Box>
                {/* Modals */}
                {showInfoModal && (
                    <PluginInfoModal
                        setShowModal={setShowInfoModal}
                        setSetting={setSelectedSetting}
                        // eslint-disable-next-line
                        setting={selectedSetting!}
                    />
                )}
                {showAddModal && (
                    <AddPluginModal
                        // eslint-disable-next-line
                        setShowModal={(toggle: boolean = false): void => {
                            setShowAddModal(toggle);
                            getSettings();
                        }}
                    />
                )}
            </Box>
        );
    }

    return <Box>configure your plugin</Box>;
};

export default PluginSettingPage;
