import React, { useState } from 'react';
import axios from 'axios';

const AddRecord = () => {
  const [form, setForm] = useState({
    PartnerId: '',
    PartnerName: '',
    CustomerId: '',
    CustomerName: '',
    // Add other fields as needed
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm({ ...form, [name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    axios.post('http://localhost:8080/write', form)
      .then(response => {
        console.log(response.data);
        // Optionally clear the form
        setForm({
          PartnerId: '',
          PartnerName: '',
          CustomerId: '',
          CustomerName: '', 
        });
      })
      .catch(error => {
        console.error('There was an error adding the record!', error);
      });
  };

  return (
    <div>
      <h2>Add Record</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>PartnerId:</label>
          <input type="text" name="PartnerId" value={form.PartnerId} onChange={handleChange} />
        </div>
        <div>
          <label>PartnerName:</label>
          <input type="text" name="PartnerName" value={form.PartnerName} onChange={handleChange} />
        </div>
        <div>
          <label>CustomerId:</label>
          <input type="text" name="CustomerId" value={form.CustomerId} onChange={handleChange} />
        </div>
        <div>
          <label>CustomerName:</label>
          <input type="text" name="CustomerName" value={form.CustomerName} onChange={handleChange} />
        </div> 
        <button type="submit">Add Record</button>
      </form>
    </div>
  );
};

export default AddRecord;
