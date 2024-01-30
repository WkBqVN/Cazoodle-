import React, { useState, useEffect } from 'react';

const DynamicForm = () => {
    const [formFields, setFormFields] = useState(() => []);
    const [loading, setLoading] = useState(true);

    const handleAddField = () => {
        const newField = { id: formFields.length + 1, title: "", value: '', type: 'text' };
        setFormFields([...formFields, newField]);
        console.log(formFields)
    };
    useEffect(() => {
        // Fetch data from the API when the component mounts
        fetchDataFromApi();
    }, []); // The empty dependency array ensures this effect runs only once when the component mounts

    const fetchDataFromApi = async () => {
        try {
            const response = await fetch('http://localhost:5000/survey/1/1');
            const data = await response.json();
            // Assuminkg the API response has a 'fields' property containing the form fields
            setFormFields(data.message || []);
            setLoading(false);
        } catch (error) {
            console.error('Error fetching data from API:', error);
            setLoading(false);
        }
    };
    const postData = async (paramId) => {
        console.log(formFields)
        const response = await fetch('http://localhost:5000/survey/1/1', {
            method: 'POST',
            body: formFields,
        });
        if (response.ok) {
            console.log('Form data sent successfully');
            // Optionally, you can reset the form or perform other actions after successful submission
        } else {
            console.error('Failed to send form data:', response.status);
        }
    }

    const handleRemoveField = (id) => {
        const updatedFields = formFields.filter((field) => field.id !== id);
        setFormFields(updatedFields);
    };

    const handleInputChange = (id, value) => {
        const updatedFields = formFields.map((field) =>
            field.id === id ? { ...field, value } : field
        );
        setFormFields(updatedFields);
    };

    const handleQuestionChange = (id, title) => {
        const updatedFields = formFields.map((field) =>
            field.id === id ? { ...field, title } : field
        );
        setFormFields(updatedFields)
    }

    const handleSelectChange = (id, selectedType) => {
        setFormFields((prevFields) =>
            prevFields.map((field) =>
                field.id === id ? { ...field, type: selectedType } : field
            )
        );
    };

    const handleCheckboxChange = (fieldId, checkboxId) => {
        setFormFields((prevFields) =>
            prevFields.map((field) =>
                field.id === fieldId
                    ? {
                        ...field,
                        value: Array.isArray(field.value)
                            ? field.value.map((checkbox) =>
                                checkbox.id === checkboxId
                                    ? { ...checkbox, checked: !checkbox.checked }
                                    : checkbox
                            )
                            : field.value,
                    }
                    : field
            )
        );
    };

    const handleSendForm = (id) => {
        postData(id)
    }
    const handleAddCheckbox = (id) => {
        const updatedFields = formFields.map((field) =>
            field.id === id && field.type === 'checkbox'
                ? {
                    ...field,
                    value: [
                        ...(Array.isArray(field.value) ? field.value : []),
                        { id: field.value.length + 1, label: 'xyz', checked: false },
                    ],
                }
                : field
        );
        setFormFields(updatedFields);
    }

    return (
        <div>
            <h2>Form</h2>
            <form>
                {formFields.map((field) => (
                    <div key={field.id} style={{ display: 'flex', alignItems: 'center' }}>
                        <input
                            type="text"
                            value={field.title}
                            onChange={(e) => handleQuestionChange(field.id, e.target.value)}
                            style={{ marginRight: '20px' }}
                        />
                        {field.type === 'text' && (
                            <div style={{ marginRight: '20px', display: 'flex', alignItems: 'center' }}>
                                <input
                                    type="text"
                                    placeholder=""
                                    value={field.value}
                                    onChange={(e) => handleInputChange(field.id, e.target.value)}
                                />
                                <select value={field.type} key={field.id}
                                    onChange={(e) => handleSelectChange(field.id, e.target.value)}>
                                    <option value="text">Text</option>
                                    <option value="checkbox">Checkbox</option>
                                    <option value="date">Date</option>
                                    <option value="int">Number</option>
                                </select>
                            </div>
                        )}
                        {field.type === 'checkbox' && Array.isArray(field.value) ? (
                            <div style={{ marginRight: '20px' }}>
                                <button type="button" onClick={(field) => handleAddCheckbox(field.id)}>+</button>
                                {field.value.map((checkbox) => (
                                    <div key={checkbox.id}>
                                        <input
                                            type="checkbox"
                                            checked={checkbox.checked}
                                            onChange={() =>
                                                handleCheckboxChange(field.id, checkbox.id)
                                            }
                                        />
                                        <span style={{ fontSize: '12px' }}>
                                            {checkbox.label}
                                        </span>
                                    </div>
                                ))}
                                <select value={field.type}
                                    onChange={(e) => handleSelectChange(field.id, e.target.value)}>
                                    <option value="text">Text</option>
                                    <option value="checkbox">Checkbox</option>
                                    <option value="date">Date</option>
                                    <option value="int">Number</option>
                                </select>
                            </div>
                        ) : null}
                        {field.type === 'int' && (
                            <div key={field.id} style={{ marginRight: '20px', display: 'flex', alignItems: 'center' }}>
                                <div key={field.id}>
                                    <input
                                        type="text"
                                        placeholder=""
                                        value={field.value}
                                        onChange={(e) => handleInputChange(field.id, e.target.value)}
                                    />
                                </div>
                                <select value={field.type}
                                    onChange={(e) => handleSelectChange(field.id, e.target.value)}>
                                    <option value="text">Text</option>
                                    <option value="checkbox">Checkbox</option>
                                    <option value="date">Date</option>
                                    <option value="int">Number</option>
                                </select>
                            </div>
                        )}
                        {field.type === 'date' && (
                            <div key={field.id} style={{ marginRight: '20px', display: 'flex', alignItems: 'center' }}>
                                <input
                                    type="date"
                                    value={field.value}
                                    onChange={(e) => handleInputChange(field.id, e.target.value)}
                                />
                                <select value={field.type}
                                    onChange={(e) => handleSelectChange(field.id, e.target.value)}>
                                    <option value="text">Text</option>
                                    <option value="checkbox">Checkbox</option>
                                    <option value="date">Date</option>
                                    <option value="int">Number</option>
                                </select>
                            </div>
                        )}
                        <button type="button" onClick={() => handleRemoveField(field.id)}>
                            Remove
                        </button>
                    </div>
                ))
                }
            </form >
            <button type="button" onClick={handleAddField}>
                Add Field
            </button>
            <button type="button" onClick={handleSendForm}>
                Send
            </button>
        </div >
    );
};

export default DynamicForm;
