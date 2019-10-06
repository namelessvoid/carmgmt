import React from 'react';
import './App.css';
import { CarListComponent } from './cars/CarsListComponent';

const App: React.FC = () => {
  return (
    <div className="App">
      <CarListComponent />
    </div>
  );
}

export default App;
