import React, { useState, useEffect, useCallback } from 'react';
import { Box, Button, FormField, Heading, Layer, TextInput, CheckBox } from 'grommet';
import { Close, Add, Subtract } from 'grommet-icons';

import { createTask } from '../../data/tasker';
import DatePicker from '../datepicker';
import TimePicker from '../timepicker';

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
    const d = new Date();
    const defaultHour = d.getHours();
    const defaultMinute = d.getMinutes();
    const [createTaskInput, setCreateTaskInput] = useState<Partial<NewTaskInput>>({
        name: '',
        schedule: d.toISOString(),
        description: '',
        executor: '',
    });
    const [time, setTime] = useState<Time>({ hour: defaultHour, minute: defaultMinute });
    const [isRepeating, setIsRepeating] = useState<boolean>(false);
    const [args, setArgs] = useState<Argument[]>([{ key: '', value: '' }]);
    const [showCalendar, setShowCalendar] = useState<boolean>(false);
    const [disableNext, setDisableNext] = useState<boolean>(true);
    /* eslint-disable-next-line @typescript-eslint/no-unused-vars  */
    const [disableCreate, setDisableCreate] = useState<boolean>(true);

    const onDateSelect = (selectedDate: any) => {
        console.log(selectedDate);
        setCreateTaskInput({ ...createTaskInput, schedule: selectedDate });
        setShowCalendar(false);
    };

    const onTimeSelect = (e: any) => {
        const { name } = e.target;
        setTime({ ...time, [name]: e.value });
    };

    const onChange = (e: any): void => {
        const key = e.target.name;
        const value = e.target.value;

        setCreateTaskInput({ ...createTaskInput, [key]: value });
    };

    const onNext = (): void => {
        setNext(true);
        const selectedMonth = new Date(createTaskInput.schedule!).getMonth() + 1;
        const selectedDayOfWeek = new Date(createTaskInput.schedule!).getDay();
        const selectedDayOfMonth = new Date(createTaskInput.schedule!).getDate();
        const cronTab = `${time.minute} ${time.hour} ${selectedDayOfMonth} ${selectedMonth} ${selectedDayOfWeek} `;
        setCreateTaskInput({ ...createTaskInput, schedule: cronTab });
    };

    const create = async (): Promise<void> => {
        const input: any = { ...createTaskInput, args: {} };
        /* eslint-disable-next-line */
        args.map((a: Argument, i: number): void => {
            if (a.key !== '' && a.value !== '') {
                input.args[a.key] = a.value;
            }
        });

        await createTask(input);
        showModal(false);
    };

    const shouldDisableNext = useCallback(() => {
        for (const key of Object.keys(createTaskInput)) {
            const k: NewTaskInputKey = key as NewTaskInputKey;
            if (createTaskInput[k]!.length === 0) {
                setDisableNext(true);
                return;
            }
        }
        setDisableNext(false);
    }, [createTaskInput]);

    useEffect(() => {
        if (args.length > 1) {
            setDisableCreate(false);
        }
        shouldDisableNext();
    }, [args.length, createTaskInput, shouldDisableNext]);

    return (
        <Layer modal onEsc={(): void => showModal()}>
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
                                <Box direction="row" gap="small">
                                    <DatePicker
                                        setShowCalendar={setShowCalendar}
                                        showCalendar={showCalendar}
                                        date={createTaskInput.schedule}
                                        onSelect={onDateSelect}
                                    />
                                    <TimePicker onChange={onTimeSelect} time={time} />
                                </Box>
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
                                onClick={(): void => onNext()}
                                style={{ borderRadius: '7px' }}
                                disabled={disableNext}
                            />
                        </Box>
                    ) : (
                        <Box fill gap="small">
                            <Heading level="4">Args: </Heading>
                            <ArgsFieldList args={args} setArgs={setArgs} />
                            <CheckBox
                                checked={isRepeating}
                                onChange={() => setIsRepeating(!isRepeating)}
                                label="Repeating?"
                            />
                            <Button
                                label="Create"
                                style={{ borderRadius: '7px' }}
                                disabled={disableCreate}
                                onClick={(): Promise<void> => create()}
                            />
                        </Box>
                    )}
                </Box>
            </Box>
        </Layer>
    );
};

export default CreateTaskModal;
