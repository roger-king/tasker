import React, { useState, FormEvent } from 'react';
import { Box, DropButton, Select, Text } from 'grommet';
import { Clock } from 'grommet-icons';

import { range } from '../../utils/array';

interface TimePickerProps {
    time: Time;
    onChange(e: FormEvent<HTMLInputElement>): void;
}

const TimePicker: React.FC<TimePickerProps> = (props: TimePickerProps) => {
    const {
        time: { hour, minute },
    } = props;
    const [showPicker, setShowPicker] = useState<boolean>(false);
    const hours = range(0, 23);
    const minutes = range(0, 59);
    const AMPM = hour >= 12 ? 'pm' : 'am';
    return (
        <DropButton
            open={showPicker}
            onClose={() => setShowPicker(false)}
            onOpen={() => setShowPicker(true)}
            dropAlign={{ top: 'bottom', left: 'left' }}
            dropContent={
                <Box direction="row" width="300px">
                    <Select options={hours} open size="xsmall" value={String(hour)} icon={false} plain />
                    <Select options={minutes} open size="xsmall" value={String(minute)} icon={false} plain />
                    <Select options={['am', 'pm']} open size="xsmall" value={AMPM} icon={false} plain />
                </Box>
            }
        >
            <Box direction="row" align="center" justify="between" pad={{ left: '10px', right: '10px' }}>
                <Text>{`${hour}: ${minute} ${'PM'}`}</Text>
                <Box alignSelf="end" margin="xsmall">
                    <Clock color="accent-1" />
                </Box>
            </Box>
        </DropButton>
    );
};

export default TimePicker;
