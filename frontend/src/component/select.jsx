// CheckboxInput.js
import React from 'react';

const SelectInput = ({ value, onChange, id }) => (
    <select value={value} key={id}
        onChange={(e) => onChange(id, e.target.value)}>
        <option value="text">Text</option>
        <option value="checkbox">Checkbox</option>
        <option value="date">Date</option>
        <option value="int">Number</option>
    </select>
);

export default SelectInput;
