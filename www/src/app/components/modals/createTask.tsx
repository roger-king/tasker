import React, { useState } from 'react';
import { Box, Button, Form, FormField, Heading, Layer } from 'grommet';
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
    const newRow = () => {
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

    return (
        <Layer modal onClickOutside={() => showModal()} onEsc={() => showModal()}>
            <Box width="medium" pad="medium">
                <Box direction="row">
                    <Button icon={<Close size="medium" />} onClick={() => showModal()} />
                    <Heading level="4">Create Task</Heading>
                </Box>
                <Form>
                    {!next ? (
                        <Box>
                            <FormField label="Name" />
                            <FormField label="Description" />
                            <FormField label="Schedule" />
                            <FormField label="Executor" />
                            <Button label="Next" onClick={() => setNext(true)} />
                        </Box>
                    ) : (
                        <Box>
                            <Heading level="4">Args: </Heading>
                            <ArgsFieldList />
                            <Button label="Create" />
                        </Box>
                    )}
                </Form>
            </Box>
        </Layer>
    );
};

export default CreateTaskModal;
