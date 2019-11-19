import React, { useState } from 'react';
import { Box, Button, FormField, Heading, Layer, TextInput } from 'grommet';
import { Close, Add, Subtract } from 'grommet-icons';

interface ArgsProps {
    keyArg: string;
    setKey: any;
    valueArg: any;
    setValue: any;
}

const ArgsField: React.FC<ArgsProps> = (props: ArgsProps): JSX.Element => {
    const { keyArg, valueArg, setKey, setValue } = props;
    const onChange = (e: any) => {
        if (e.target.name === 'key') {
            setKey(e.target.value);
        } else if (e.target.name === 'value') {
            setValue(e.target.value);
        }
    };
    return (
        <Box direction="row" gap="small">
            <TextInput name="key" placeholder="Key" onChange={onChange} value={keyArg} />
            <TextInput name="value" placeholder="Value" onChange={onChange} value={valueArg} />
        </Box>
    );
};

interface ArgsListProps {
    args: Map<string, string | Record<string, any>>;
    setArgs: any;
}

const ArgsFieldList: React.FC<ArgsListProps> = (props: ArgsListProps) => {
    const { args, setArgs } = props;
    const [keyArg, setKey] = useState<string>('');
    const [valueArg, setValue] = useState<any>('');
    const [fields, setFields] = useState<number[]>([1]);
    const newRow = (): void => {
        setFields([...fields, fields.length + 1]);
        setArgs(args.set(keyArg, valueArg));
    };

    return (
        <Box direction="column">
            {fields.map((i: number) => {
                if (i === fields.length) {
                    return (
                        <Box key={i} direction="row">
                            <ArgsField keyArg={keyArg} setKey={setKey} valueArg={valueArg} setValue={setValue} />
                            <Button icon={<Add size="small" />} onClick={newRow} />
                        </Box>
                    );
                }
                return (
                    <Box key={i} direction="row">
                        <ArgsField keyArg={keyArg} setKey={setKey} valueArg={valueArg} setValue={setValue} />
                        <Button icon={<Subtract size="small" />} onClick={() => console.log('hello')} />
                    </Box>
                );
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
    const [args, setArgs] = useState<Map<string, string | Record<string, any>>>(new Map());
    const [disableNext, setDisableNext] = useState<boolean>(true);
    /* eslint-disable-next-line @typescript-eslint/no-unused-vars  */
    const [disableCreate, setDisableCreate] = useState<boolean>(true);

    const onChange = (e: any): void => {
        const key = e.target.name;
        const value = e.target.value;

        if (value.length > 0) {
            setCreateTaskInput(createTaskInput.set(key, value));
        } else {
            createTaskInput.delete(key);
            setCreateTaskInput(createTaskInput);
        }

        if (createTaskInput.size === 4) {
            setDisableNext(false);
        } else {
            setDisableNext(true);
        }
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
                                disabled={!disableNext}
                            />
                        </Box>
                    ) : (
                        <Box fill>
                            <Heading level="4">Args: </Heading>
                            <ArgsFieldList args={args} setArgs={setArgs} />
                            <Button label="Create" style={{ borderRadius: '7px' }} disabled={disableCreate} />
                        </Box>
                    )}
                </Box>
            </Box>
        </Layer>
    );
};

export default CreateTaskModal;
