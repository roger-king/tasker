import React from 'react';
import { Box, DropButton, Calendar, Text } from 'grommet';
import { Calendar as CalendarIcon } from 'grommet-icons';

interface DatePickerProps {
    showCalendar: boolean;
    setShowCalendar: any;
    date: any;
    onSelect: any;
}

const DatePicker: React.FC<DatePickerProps> = (props: DatePickerProps) => {
    const { showCalendar, setShowCalendar, date, onSelect } = props;
    return (
        <DropButton
            open={showCalendar}
            onClose={() => setShowCalendar(false)}
            onOpen={() => setShowCalendar(true)}
            dropAlign={{ top: 'bottom', right: 'right' }}
            dropContent={<Calendar date={date} onSelect={onSelect} size="small" />}
        >
            <Box direction="row" align="center" justify="between" pad={{ left: '10px', right: '10px' }}>
                <Text>{date ? new Date(date).toLocaleDateString() : new Date().toLocaleDateString()}</Text>
                <Box alignSelf="end" margin="xsmall">
                    <CalendarIcon color="accent-1" />
                </Box>
            </Box>
        </DropButton>
    );
};

export default DatePicker;
