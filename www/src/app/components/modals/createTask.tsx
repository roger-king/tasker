import React, { useState, useEffect } from 'react';
import { Box, Button, FormField, Heading, Layer, TextInput } from 'grommet';
import { Close, Add, Subtract } from 'grommet-icons';

import { createTask } from '../../data/tasker';

interface ArgsProps {
    index: number;
    addRow(key: string, value: any): void;
    isNewRow: boolean;
    removeRow(index: number, key: string): void;
}

const ArgsField: React.FC<ArgsProps> = (props: ArgsProps): JSX.Element => {
    const { addRow, isNewRow, index, removeRow } = props;
    const [disableAddRow, setdisableAddRow] = useState<boolean>(true);
    const [currKey, setCurrKey] = useState<string>('');
    const [currValue, setCurrValue] = useState<any>('');

    const onChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
        const { name, value } = e.target;
        if (name === 'key') {
            setCurrKey(value);
        } else if (name === 'value') {
            setCurrValue(value);
        }

        if (currKey.length > 0 && currValue.length > 0) {
            setdisableAddRow(false);
        } else {
            setdisableAddRow(true);
        }
    };

    return (
        <Box direction="row" gap="small">
            <TextInput name="key" placeholder="Key" onChange={onChange} value={currKey} />
            <TextInput name="value" placeholder="Value" onChange={onChange} value={currValue} />
            {isNewRow ? (
                <Button
                    icon={<Add size="small" />}
                    onClick={(): void => addRow(currKey, currValue)}
                    disabled={disableAddRow}
                />
            ) : (
                <Button icon={<Subtract size="small" />} onClick={(): void => removeRow(index, currKey)} />
            )}
        </Box>
    );
};

interface ArgsListProps {
    args: Argument[];
    setArgs: any;
}

const ArgsFieldList: React.FC<ArgsListProps> = (props: ArgsListProps) => {
    const { args, setArgs } = props;
    const [fields, setFields] = useState<number[]>([Number(1)]);

    const addRow = (key: string, value: any): void => {
        const foundArg = args.filter((arg: Argument) => arg.key === key);

        if (foundArg.length === 0) {
            setFields([...fields, fields.length + 1]);
            setArgs([...args, { key, value }]);
        }
    };

    const removeRow = (index: number, key: string): void => {
        const f = fields.map(num => {
            if (num !== index) {
                return num;
            }
            return 0;
        });
        setFields(f);

        const newArgs = args.filter((arg: Argument) => arg.key !== key);
        setArgs([...newArgs]);
    };

    return (
        <Box direction="column" gap="small">
            {/* eslint-disable react/no-array-index-key */}
            {/* eslint-disable-next-line */}
            {fields.map((num: number, i: number): JSX.Element | void => {
                if (num !== 0) {
                    return (
                        <ArgsField
                            key={i}
                            index={num}
                            addRow={addRow}
                            isNewRow={num === fields[fields.length - 1]}
                            removeRow={removeRow}
                        />
                    );
                }
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
    const [createTaskInput, setCreateTaskInput] = useState<Partial<NewTaskInput>>({
        name: '',
        schedule: '',
        description: '',
        executor: '',
    });
    const [args, setArgs] = useState<Argument[]>([{ key: '', value: '' }]);
    const [disableNext, setDisableNext] = useState<boolean>(true);
    /* eslint-disable-next-line @typescript-eslint/no-unused-vars  */
    const [disableCreate, setDisableCreate] = useState<boolean>(true);

    const onChange = (e: any): void => {
        const key = e.target.name;
        const value = e.target.value;

        setCreateTaskInput({ ...createTaskInput, [key]: value });

        if (Object.keys(createTaskInput).length === 4) {
            setDisableNext(false);
        } else {
            setDisableNext(true);
        }
    };

    const create = async (): Promise<void> => {
        const input: any = { ...createTaskInput, args: {} };
        /* eslint-disable-next-line */
        args.map((a: Argument, i: number): void => {
            input.args[a.key] = a.value;
        });

        await createTask(input);
    };

    useEffect(() => {
        if (args.length > 1) {
            setDisableCreate(false);
        }
    }, [args.length]);

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
                                <TextInput name="name" onChange={onChange} value={createTaskInput.name} required />
                            </FormField>
                            <FormField label="Description">
                                <TextInput
                                    name="description"
                                    onChange={onChange}
                                    value={createTaskInput.description}
                                    required
                                />
                            </FormField>
                            <FormField label="Schedule">
                                <TextInput
                                    name="schedule"
                                    onChange={onChange}
                                    value={createTaskInput.schedule}
                                    required
                                />
                            </FormField>
                            <FormField label="Executor">
                                <TextInput
                                    name="executor"
                                    onChange={onChange}
                                    value={createTaskInput.executor}
                                    required
                                />
                            </FormField>
                            <Button
                                label="Next"
                                onClick={(): void => setNext(true)}
                                style={{ borderRadius: '7px' }}
                                disabled={disableNext}
                            />
                        </Box>
                    ) : (
                        <Box fill gap="small">
                            <Heading level="4">Args: </Heading>
                            <ArgsFieldList args={args} setArgs={setArgs} />
                            <Button
                                label="Create"
                                style={{ borderRadius: '7px' }}
                                disabled={disableCreate}
                                onClick={(): void => create()}
                            />
                        </Box>
                    )}
                </Box>
            </Box>
        </Layer>
    );
};

export default CreateTaskModal;
