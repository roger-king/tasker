import React, { useState, useEffect } from 'react';
import {
    Box,
    Heading,
    Text,
    CheckBox,
    Table,
    TableRow,
    TableCell,
    TableBody,
    Button,
    Anchor,
    Form,
    FormField,
    TextInput,
} from 'grommet';
import { StatusInfo } from 'grommet-icons';
import { listSettings, toggleActivePluginSetting, createSetting } from '../../data/settings';
import Modal, { BaseModalProps } from '../../components/modals';

interface AddPluginModalProps extends BaseModalProps {
    setShowModal: any;
}

const AddPluginModal: React.FC<AddPluginModalProps> = (props: AddPluginModalProps): JSX.Element => {
    const { setShowModal } = props;
    return (
        <Modal setShowModal={setShowModal} header="Add Repository" width="large" onClickOutside onEsc>
            <Box>
                <Form
                    onSubmit={async ({ value }: any): Promise<void> => {
                        const input = { ...value, type: 'plugin' };
                        await createSetting(input);
                        setShowModal(false);
                    }}
                >
                    <FormField
                        name="repo_name"
                        component={TextInput}
                        label="Github Repository"
                        required
                        placeholder="<github_username>/<repository_name>"
                    />
                    <FormField name="description" component={TextInput} label="Description" />
                    <FormField
                        name="build_folder_name"
                        component={TextInput}
                        placeholder="build"
                        label="Build Folder Name"
                        help="name of the folder your Go build process places your plugins"
                        required
                    />
                    <FormField name="active" component={CheckBox} pad label="Is Active?" checked required />
                    <Button label="add" primary type="submit" />
                </Form>
            </Box>
        </Modal>
    );
};

interface PluginInfoModalProps extends BaseModalProps {
    setSetting: any;
    setting: Setting;
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
                <Heading level="4">{setting.repo_name}</Heading>
                <Text>{setting.description}</Text>
                <Anchor href={`https://github.com/${setting.repo_name}`} target="_blank">
                    View Github
                </Anchor>
            </Box>
        </Modal>
    );
};

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
