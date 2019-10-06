import React, { Component } from 'react';

interface Car {
    ID: number;
    Name: string;
}

interface CarListComponentState {
    cars: Car[];
}

export class CarListComponent extends Component<{}, CarListComponentState> {
    constructor(props: any) {
        super(props)
        this.state = {
            cars: []
        };
    }
    componentDidMount() {
        console.log('Fetch cars')
        fetch('http://localhost:8080/cars')
            .then(response => response.json())
            .then(cars => this.setState({ cars }))
    }

    render() {
        return (
            <div>
                <h2>Cars</h2>
                {this.state.cars.map(car => {
                    console.log(car);
                    return <div key={car.ID}>{car.Name}</div>;
                })}
            </div>
        )
    }
}