import React from 'react';
import styled from 'styled-components';
import { Box, Button, CheckBox, Heading, Text } from 'grommet';
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
            <Button icon={<MoreVertical size="medium" />} />
        </Box>
    );
};

export default styled(Task)``;
