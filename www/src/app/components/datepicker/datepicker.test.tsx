import React from 'react';
import renderer from 'react-test-renderer';
import DatePicker from './index';

describe('<DatePicker />', () => {
    it('Matches the snapshot with datepicker closed', () => {
        const onSelect = jest.fn();
        const setShowCalendar = jest.fn();
        const datepicker = renderer
            .create(
                <DatePicker date={null} onSelect={onSelect} setShowCalendar={setShowCalendar} showCalendar={false} />,
            )
            .toJSON();
        expect(datepicker).toMatchSnapshot();
    });

    it('Matches the snapshot with datepicker open', () => {
        const onSelect = jest.fn();
        const setShowCalendar = jest.fn();
        const datepicker = renderer
            .create(<DatePicker date={null} onSelect={onSelect} setShowCalendar={setShowCalendar} showCalendar />)
            .toJSON();
        expect(datepicker).toMatchSnapshot();
    });
});
