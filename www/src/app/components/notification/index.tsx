import React from 'react';
import { Layer, Box, Heading, Text } from 'grommet';
import styled from 'styled-components';

import { StatusCritical, Flag, Validate, Task } from 'grommet-icons';

interface NotificationProps {
    className?: string;
    msg: TaskerNotification;
}

const Notification: React.FC<NotificationProps> = (props: NotificationProps): JSX.Element => {
    const { className, msg } = props;

    const renderIcon = (): JSX.Element => {
        // eslint-disable-next-line
        const { type } = msg;
        if (type === 'success') {
            return <Validate color="status-ok" size="large" />;
        }
        if (type === 'warn') {
            return <Flag color="status-warning" size="large" />;
        }
        if (type === 'fail') {
            return <StatusCritical color="status-critical" size="large" />;
        }
        return <Task color="brand" size="large" />;
    };

    return (
        <Layer className={className} modal={false} animation="slide" position="top-right">
            <Box align="start" pad="small" direction="row" gap="medium">
                {renderIcon()}
                <Box align="start">
                    <Heading>{msg.title}</Heading>
                    <Text>{msg.body}</Text>
                </Box>
            </Box>
        </Layer>
    );
};

export default styled(Notification)`
    width: 300px;
    min-width: 200px;
    height: 80px;
    min-height: 50px;
    margin-top: 90px;
    margin-right: 20px;
`;
