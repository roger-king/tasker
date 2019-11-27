import React, { useState, FormEvent } from 'react';
import { Box, DropButton, Select, Text } from 'grommet';
import { Clock } from 'grommet-icons';

import { range } from '../../utils/array';

interface TimePickerProps {
    time: Time;
    onChange(e: FormEvent<HTMLInputElement>): void;
}

const TimePicker: React.FC<TimePickerProps> = (props: TimePickerProps): JSX.Element => {
    const {
        onChange,
        time: { hour, minute },
    } = props;
    const [showPicker, setShowPicker] = useState<boolean>(false);
    // TODO: fix this to display leading 0
    const defaultHour = String(hour).length === 1 ? `0${String(hour)}` : String(hour);
    const defaultMinute = String(minute).length === 1 ? `0${String(minute)}` : String(minute);
    const hours = range(0, 23);
    const minutes = range(0, 59);
    // TODO: support 12 hour time
    const AMPM = hour >= 12 ? 'pm' : 'am';
    return (
        <DropButton
            open={showPicker}
            onClose={(): void => setShowPicker(false)}
            onOpen={(): void => setShowPicker(true)}
            dropAlign={{ top: 'bottom', left: 'left' }}
            dropContent={
                <Box direction="row" width="300px">
                    <Select
                        name="hour"
                        options={hours}
                        open
                        size="xsmall"
                        value={defaultHour}
                        icon={false}
                        plain
                        onChange={onChange}
                    />
                    <Select
                        name="minute"
                        options={minutes}
                        open
                        size="xsmall"
                        value={defaultMinute}
                        icon={false}
                        plain
                        onChange={onChange}
                    />
                </Box>
            }
        >
            <Box direction="row" align="center" justify="between" pad={{ left: '10px', right: '10px' }}>
                <Text>{`${hour}: ${minute} ${AMPM}`}</Text>
                <Box alignSelf="end" margin="xsmall">
                    <Clock color="accent-1" />
                </Box>
            </Box>
        </DropButton>
    );
};

export default TimePicker;
