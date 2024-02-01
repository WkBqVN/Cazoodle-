import React, { useState } from 'react';
import DynamicForm from './dynamic_form';
import GetForm from './getForm';

const NewForm = () => {
    const [showDynamicForm, setShowDynamicForm] = useState(false);
    const [showGetForm, setShowGetForm] = useState(false);

    const handleCreateForm = () => {
        setShowDynamicForm(true);
        setShowGetForm(false)
    };

    const handleGetForm = () => {
        setShowDynamicForm(false)
        setShowGetForm(true)
    };


    return (
        <div>
            {showDynamicForm ? (
                <DynamicForm />
            ) : (
                <div>
                    <button type="button" onClick={handleCreateForm}>Create Form</button>
                </div>
            )}
            {showGetForm ? (
                <GetForm />
            ) : (
                <div>
                    <button type="button" onClick={handleGetForm}>Get Form</button>
                </div>
            )}
        </div>
    );
};

export default NewForm;