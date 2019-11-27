import React from 'react';
import { Box, Button, Form, FormField, CheckBox, TextInput } from 'grommet';

import Modal, { BaseModalProps } from './index';
import { createSetting } from '../../data/settings';

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

export default AddPluginModal;
