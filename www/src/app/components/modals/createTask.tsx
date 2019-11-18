import React, { useState } from 'react';
import { Box, Button, FormField, Heading, Layer, TextInput } from 'grommet';
import { Close, Add } from 'grommet-icons';

const ArgsField: React.FC<{}> = (): JSX.Element => {
    return (
        <Box direction="row" gap="small">
            <FormField label="Key" />
            <FormField label="Value" />
        </Box>
    );
};

const ArgsFieldList: React.FC<{}> = () => {
    const [fields, setFields] = useState<number[]>([1]);
    const newRow = (): void => {
        setFields([...fields, fields.length + 1]);
    };

    return (
        <Box direction="column">
            {fields.map((i: number) => {
                console.log(i === fields.length);
                if (i === fields.length) {
                    return (
                        <Box key={i} direction="row">
                            <ArgsField key={i} />
                            <Button icon={<Add size="small" />} onClick={newRow} />
                        </Box>
                    );
                }
                return <ArgsField key={i} />;
            })}
        </Box>
    );
};

interface CreateTaskModalProps {
    showModal: any;
}

const CreateTaskModal: React.FC<CreateTaskModalProps> = (props: CreateTaskModalProps): JSX.Element => {
    const { showModal } = props;
    const [next, setNext] = useState<boolean>(false);
    const [createTaskInput, setCreateTaskInput] = useState<Map<string, string | Record<string, any>>>(new Map());
    /* eslint-disable @typescript-eslint/no-unused-vars */
    const [disableNext, _setDisableNext] = useState<boolean>(true);

    const onChange = (e: any): void => {
        const key = e.target.name;
        const value = e.target.value;

        setCreateTaskInput(createTaskInput.set(key, value));

        console.log(createTaskInput.size);
    };

    return (
        <Layer modal onClickOutside={(): void => showModal()} onEsc={(): void => showModal()}>
            <Box width="large" pad="medium">
                <Box direction="row">
                    <Button icon={<Close size="medium" />} onClick={(): void => showModal()} />
                    <Heading level="4">Create Task</Heading>
                </Box>
                <Box>
                    {!next ? (
                        <Box>
                            <FormField label="Name">
                                <TextInput name="name" onChange={onChange} required />
                            </FormField>
                            <FormField label="Description">
                                <TextInput name="description" onChange={onChange} required />
                            </FormField>
                            <FormField label="Schedule">
                                <TextInput name="schedule" onChange={onChange} required />
                            </FormField>
                            <FormField label="Executor">
                                <TextInput name="executor" onChange={onChange} required />
                            </FormField>
                            <Button
                                label="Next"
                                onClick={(): void => setNext(true)}
                                style={{ borderRadius: '7px' }}
                                disabled={disableNext}
                            />
                        </Box>
                    ) : (
                        <Box fill>
                            <Heading level="4">Args: </Heading>
                            <ArgsFieldList />
                            <Button label="Create" />
                        </Box>
                    )}
                </Box>
            </Box>
        </Layer>
    );
};

export default CreateTaskModal;
