import React, { useState } from 'react';
import { GoogleMap, Marker, useLoadScript } from '@react-google-maps/api';

const containerStyle = {
    width: '100%',
    height: '400px',
};

const center = {
    lat: -3.745,
    lng: -38.523,
};

const LocationPicker: React.FC = () => {
    // States to hold the selected position (lat/lng)
    const [selectedPosition, setSelectedPosition] = useState<{ lat: number, lng: number } | null>(null);

    // Load the Google Maps script
    const { isLoaded } = useLoadScript({
        googleMapsApiKey: process.env.REACT_APP_GOOGLE_MAPS_API_KEY || '', // Use your API key here
    });

    // Function to handle map click and update the selected position
    const handleMapClick = (event: google.maps.MapMouseEvent) => {
        if (event.latLng) {
            const lat = event.latLng.lat();
            const lng = event.latLng.lng();
            setSelectedPosition({ lat, lng });
        }
    };

    // Show loading message while map is loading
    if (!isLoaded) return <div>Loading...</div>;

    return (
        <div>
            <h3>Pick a Location</h3>
            <GoogleMap
                mapContainerStyle={containerStyle}
                center={center}
                zoom={10}
                onClick={handleMapClick}
            >
                {selectedPosition && (
                    <Marker position={selectedPosition} />
                )}
            </GoogleMap>

            {/* Display selected position if available */}
            {selectedPosition && (
                <div>
                    <p>Selected Latitude: {selectedPosition.lat}</p>
                    <p>Selected Longitude: {selectedPosition.lng}</p>
                </div>
            )}
        </div>
    );
};

export default LocationPicker;
