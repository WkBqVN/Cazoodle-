import { useState } from "react";
import SelectInput from "./select";
import DynamicForm from "./dynamic_form";

const GetForm = () => {
    const [formFields, setFormFields] = useState(() => []);
    const [formId, setFormId] = useState("");
    const [form, setForm] = useState(false)

    const fetchDataFromApi = async (id) => {
        if (id == null) {
            id = 0
        }
        try {
            const response = await fetch(`http://localhost:5000/survey/forms/${id}`);
            const data = await response.json();
            // Assuminkg the API response has a 'fields' property containing the form fields
            setFormFields(data.message || []);
        } catch (error) {
            console.error('Error fetching data from API:', error);
        }
    };

    const handleGetForm = async () => {
        try {
            await fetchDataFromApi(formId)
            setForm(true)
        } catch (error) {
            console.error("err here")
        }
    }
    const postForm = async (clientId, surveyId) => {
        const response = await fetch(`http://localhost:5000/survey/${clientId}/${surveyId}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formFields),
        });
        if (response.ok) {
            console.log('Form data sent successfully');
        } else {
            console.error('Failed to send form data:', response.status);
        }
    }
    const handleSendForm = async () => {
        try {
            let clientId = 1
            let surveyId = 1
            return postForm(clientId, surveyId)
        } catch (error) {
            console.log(error("err here"))
        }
    }

    return (
        <div>
            <h2>GET FORM</h2>
            <input
                type="text"
                placeholder='form'
                value={formId}
                onChange={(e) => setFormId(e.target.value)}
            />
            <button type="button" onClick={handleGetForm}>
                GET FORM
            </button>
            <button type="button" onClick={handleSendForm}>
                SEND FORM
            </button>
            {form ? (
                <DynamicForm formFields={formFields} />
            ) : (
                <div>
                    <button type="button" >Done</button>
                </div>
            )}
        </div >
    );
};

export default GetForm;