import React from 'react';
import styled from 'styled-components';
import { Box, Button, DropButton, CheckBox, Heading, Text } from 'grommet';
import { MoreVertical } from 'grommet-icons';

interface TaskProps {
    className?: string;
    name: string;
    description: string;
    enabled: boolean;
    complete: boolean;
    runTime: string;
}
const Task: React.FC<TaskProps> = (props: TaskProps): JSX.Element => {
    const { className, name, description, enabled, complete } = props;
    const completeText = complete ? 'Completed' : 'Not Complete';
    return (
        <Box
            className={className}
            direction="row-responsive"
            width="100vw"
            border
            align="center"
            pad={{ left: '60px' }}
            height="75px"
            gap="small"
        >
            <CheckBox disabled value="" checked={enabled} />
            <Box width="20%" align="start">
                <Heading level="4" margin="small">
                    {name}
                </Heading>
            </Box>
            <Box width="50%" align="center">
                <Text wordBreak="break-all">{description}</Text>
            </Box>
            <Box width="30%" justify="end" direction="row" align="center">
                <Text>
                    <i>{completeText}</i>
                </Text>
                <DropButton
                    icon={<MoreVertical size="medium" />}
                    dropContent={
                        <Box align="start" width="120px" gap="small">
                            <Button plain label="View" style={{ width: '100%', padding: '10px' }} hoverIndicator />
                            <Button plain label="Edit" style={{ width: '100%', padding: '10px' }} hoverIndicator />
                            <Button plain label="Delete" style={{ width: '100%', padding: '10px' }} hoverIndicator />
                        </Box>
                    }
                    dropAlign={{ top: 'bottom' }}
                    alignSelf="end"
                />
            </Box>
        </Box>
    );
};

export default styled(Task)``;
