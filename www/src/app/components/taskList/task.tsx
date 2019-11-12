import React from 'react';
import styled from 'styled-components';
import { Box, Button, DropButton, CheckBox, Heading, Text } from 'grommet';
import { MoreVertical } from 'grommet-icons';

interface TaskProps {
    className?: string;
    name: string;
    description: string;
    isSet: boolean;
    complete: boolean;
    runTime: string;
}
const Task: React.FC<TaskProps> = (props: TaskProps): JSX.Element => {
    const { className, name, description, isSet, complete } = props;
    const completeText = complete ? 'Completed' : 'Not Complete';
    return (
        <Box
            className={className}
            direction="row"
            width="100vw"
            border
            align="center"
            justify="between"
            pad={{ left: '60px' }}
            height="75px"
            gap="small"
        >
            <CheckBox disabled value="" checked={isSet} />
            <Heading level="4" margin="small">
                {name}
            </Heading>
            <Text>{description}</Text>
            <Text>
                <i>{completeText}</i>
            </Text>
            <DropButton
                icon={<MoreVertical size="medium" />}
                dropContent={
                    <Box align="start" width="120px" gap="small" pad="small">
                        <Button plain label="View" style={{ width: '100%' }} />
                        <Button plain label="Edit" style={{ width: '100%' }} />
                        <Button plain label="Delete" style={{ width: '100%' }} />
                    </Box>
                }
                dropAlign={{ top: 'bottom' }}
            />
        </Box>
    );
};

export default styled(Task)``;
