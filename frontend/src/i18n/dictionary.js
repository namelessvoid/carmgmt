import { dictionary } from 'svelte-i18n';

dictionary.set({
    'de-DE': {
        'fleet': {
            'backToOverview': 'Zurück zur Flottenübersicht',
            'overview': 'Flottenübersicht',
            'vehicle' : {
                'id': '#',
                'name': 'Name',
                'refuellings': 'Tankvorgänge',
                'validation': {
                    'nameNotEmpty': 'Der Name darf nicht leer sein'
                }
            },
            'refuelling': {
                'time': 'Datum',
                'price': 'Rechnungsbetrag',
                'amount': 'Abgabe',
                'pricePerLiter': 'Preis pro Liter',
                'tripKilometers': 'Tripkilometer',
                'consumption': 'Verbrauch',
                'emptyList': 'Es sind noch keine Tankvorgänge vorhanden.'
            },
            'addVehicle': 'Fahrzeug hinzufügen',
            'vehicleAdded': 'Fahrzeug wurde hinzugefügt.',
            'addRefuelling': 'Tankvorgang hinzufügen'
        },
        'auth': {
            'username': 'Benutzername',
            'password': 'Passwort',
            'login': 'Anmelden'
        },
        'loading': 'Lädt...',
        'error': {
            'networkFailure': 'Sie sind gerade offline oder der Server ist gerade nicht erreichbar.',
            'unknown': 'Es ist ein Fehler aufgetreten.',
            'invalidJson': 'Es ist ein Fehler aufgetreten (Server antwortete: Ungültiges JSON).'
        }
    },
    'en': {
        'fleet': {
            'backToOverview': 'Back to fleet overview',
            'overview': 'Fleet overview',
            'vehicle': {
                'id': '#',
                'name': 'Name',
                'refuellings': 'Refuellings',
                'validation': {
                    'nameNotEmpty': 'Name must not be empty'
                }
            },
            'addVehicle': 'Add vehicle',
            'vehicleAdded': 'Vehicle has been added.',
            'refuelling': {
                'time': 'Date',
                'price': 'Price',
                'amount': 'Amount',
                'pricePerLiter': 'Price per liter',
                'tripKilometers': 'Trip kilometers',
                'consumption': 'Consumption',
                'emptyList': 'No refuellings added yet'
            },
            'addRefuelling': 'Add refuelling'
        },
        'auth': {
            'username': 'Username',
            'password': 'Password',
            'login': 'Login'
        },
        'loading': 'Loading...',
        'error': {
            'networkFailure': 'You are offline or server is not reachable at the moment.',
            'unknown': 'An error occured.',
            'invalidJson': 'An error occured (server says: invalid json).'
        }
    }
})