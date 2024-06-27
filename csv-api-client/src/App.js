import React from 'react';
import './App.css';
import RecordList from './components/RecordList';
import AddRecord from './components/AddRecord';
import EditRecord from './components/EditRecord';

function App() {
  return (
    <div className="App">
      <h1>CSV API Client</h1>
      <RecordList />
      <AddRecord />
      <EditRecord />
    </div>
  );
}

export default App;

