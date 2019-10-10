import { dictionary } from 'svelte-i18n';

dictionary.set({
    'de-DE': {
        'fleet': {
            'backToOverview': 'Zurück zur Flottenübersicht',
            'overview': 'Flottenübersicht',
            'vehicle' : {
                'id': '#',
                'name': 'Name',
            },
            'addVehicle': 'Fahrzeug hinzufügen',
            'vehicleAdded': 'Fahrzeug wurde hinzugefügt.'
        },
        'loading': 'Lädt...',
        'error': {
            'networkFailure': 'Sie sind gerade offline oder der Server ist gerade nicht erreichbar.',
            'unknown': 'Es ist ein Fehler aufgetreten.',
            'invalidJson': 'Es ist ein Fehler aufgetreten (Server antwortete mit ungültigem JSON).'
        }
    },
    'en': {
        'fleet': {
            'backToOverview': 'Back to fleet overview',
            'overview': 'Fleet overview',
            'vehicle': {
                'id': '#',
                'name': 'Name'
            },
            'addVehicle': 'Add vehicle',
            'vehicleAdded': 'Vehicle has been added.'
        },
        'loading': 'Loading...',
        'error': {
            'networkFailure': 'You are offline or server is not reachable at the moment.',
            'unknown': 'An error occured.',
            'invalidJson': 'An error occured (server responded with invalid json).'
        }
    }
})