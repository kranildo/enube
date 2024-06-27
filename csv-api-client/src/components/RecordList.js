import React, { useState, useEffect } from 'react';
import axios from 'axios';

const RecordList = () => {
  const [records, setRecords] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8080/read')
      .then(response => {
        setRecords(response.data);
      })
      .catch(error => {
        console.error('There was an error fetching the records!', error);
      });
  }, []);

  return (
    <div>
      <h2>Record List</h2>
      <table>
        <thead>
          <tr>
            <th>PartnerId</th>
            <th>PartnerName</th>
            <th>CustomerId</th>
            <th>CustomerName</th> 
          </tr>
        </thead>
        <tbody>
          {records.map((record, index) => (
            <tr key={index}>
              <td>{record.PartnerId}</td>
              <td>{record.PartnerName}</td>
              <td>{record.CustomerId}</td>
              <td>{record.CustomerName}</td> 
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default RecordList;
