import React from 'react';
import { Cycle, Trigger } from 'grommet-icons';

interface TaskTypeIconProps {
    repeat: boolean;
    size?: 'small' | 'medium' | 'large' | 'xlarge' | string;
}

const TaskTypeIcon: React.FC<TaskTypeIconProps> = (props: TaskTypeIconProps): JSX.Element => {
    const { repeat, size } = props;

    if (repeat) {
        return <Cycle size={size} color="accent-1" />;
    }

    return <Trigger size={size} color="accent-1" />;
};

export default TaskTypeIcon;
