import React, { useState } from 'react';
import { Box, TextInput } from 'grommet';
import styled from 'styled-components';

interface CronInputProps {
    className?: string;
}

const CronInput: React.FC<CronInputProps> = (props: CronInputProps): JSX.Element => {
    const { className } = props;
    const [cronTab, setCronTab] = useState<CronTab>({
        second: '*',
        minute: '*',
        hour: '*',
        dayOfMonth: '*',
        month: '*',
        dayOfWeek: '*',
    });

    const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setCronTab({ ...cronTab, [name]: value });
    };

    console.log(cronTab);

    return (
        <Box className={className} direction="row" gap="small">
            <Box className="cron-box">
                <TextInput name="second" min="0" max="59" onChange={onChange} value={cronTab.second} size="xsmall" />
            </Box>
            <Box className="cron-box">
                <TextInput name="minute" min="0" max="59" onChange={onChange} value={cronTab.minute} size="xsmall" />
            </Box>
            <Box className="cron-box">
                <TextInput name="hour" min="0" max="23" onChange={onChange} value={cronTab.hour} size="xsmall" />
            </Box>
            <Box className="cron-box">
                <TextInput
                    name="dayOfMonth"
                    min="1"
                    max="31"
                    onChange={onChange}
                    value={cronTab.dayOfMonth}
                    size="xsmall"
                />
            </Box>
            <Box className="cron-box">
                <TextInput name="month" min="1" max="12" onChange={onChange} value={cronTab.month} size="xsmall" />
            </Box>
            <Box className="cron-box">
                <TextInput
                    name="dayOfWeek"
                    min="0"
                    max="7"
                    onChange={onChange}
                    value={cronTab.dayOfWeek}
                    size="xsmall"
                />
            </Box>
        </Box>
    );
};

export default styled(CronInput)``;
