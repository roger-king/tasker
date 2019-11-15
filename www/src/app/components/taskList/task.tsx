import React, { useState } from 'react';
import styled from 'styled-components';
import { Box, Button, DropButton, CheckBox, Heading, Text } from 'grommet';
import { MoreVertical } from 'grommet-icons';
import { useHistory } from 'react-router';
import DeleteTaskModal from '../modals/delete';

interface TaskProps {
    className?: string;
    id: string;
    name: string;
    description: string;
    enabled: boolean;
    complete: boolean;
    runTime: string;
}
const Task: React.FC<TaskProps> = (props: TaskProps): JSX.Element => {
    const { className, name, description, enabled, complete, id } = props;
    const completeText = complete ? 'Completed' : 'Not Executed';
    const [showModal, setShowModal] = useState<boolean>(false);
    const history = useHistory();

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
                            <Button
                                plain
                                label="View"
                                style={{ width: '100%', padding: '10px' }}
                                hoverIndicator
                                onClick={(): void => history.push(`/tasker/admin/task/${id}`)}
                            />
                            <Button
                                plain
                                label="Delete"
                                style={{ width: '100%', padding: '10px' }}
                                hoverIndicator
                                onClick={() => setShowModal(true)}
                            />
                        </Box>
                    }
                    dropAlign={{ top: 'bottom' }}
                    alignSelf="end"
                />
            </Box>
            {showModal && <DeleteTaskModal showModal={setShowModal} taskName={name} taskId={id} />}
        </Box>
    );
};

export default styled(Task)``;
