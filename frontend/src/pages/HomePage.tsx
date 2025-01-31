import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import classes from './HomePage.module.scss';

interface Party {
  id: string;
  title: string;
  date: string;
  location: string;
  description: string;
  attendees: number;
  imageUrl?: string;
}

export const Homepage = () => {
  const navigate = useNavigate();
  const [localParties, setLocalParties] = useState<Party[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    // TODO: Replace with actual API call
    const fetchLocalParties = async () => {
      // Simulated API response
      const mockParties: Party[] = [
        {
          id: '1',
          title: 'Summer Beach Party',
          date: '2024-07-15',
          location: 'Miami Beach',
          description: 'Join us for a fantastic beach party with live music and great vibes!',
          attendees: 45,
          imageUrl: 'https://images.unsplash.com/photo-1533174072545-7a4b6ad7a6c3',
        },
        {
          id: '2',
          title: 'Rooftop Jazz Night',
          date: '2024-06-20',
          location: 'Downtown Skybar',
          description: 'An evening of smooth jazz under the stars.',
          attendees: 30,
          imageUrl: 'https://images.unsplash.com/photo-1516450360452-9312f5e86fc7',
        },
      ];

      setLocalParties(mockParties);
      setLoading(false);
    };

    fetchLocalParties();
  }, []);

  return (
    <div className={classes.homepage}>
      <section className={classes.hero}>
        <div className={classes.heroContent}>
          <h1>Welcome to Party Organizer</h1>
          <p>Discover, create, and join amazing events in your area</p>
          <div className={classes.heroButtons}>
            <button 
              className={classes.primaryButton}
              onClick={() => navigate('/register')}
            >
              Get Started
            </button>
            <button 
              className={classes.secondaryButton}
              onClick={() => navigate('/overview/discover')}
            >
              Explore Events
            </button>
          </div>
        </div>
      </section>

      <section className={classes.features}>
        <h2>Why Choose Party Organizer?</h2>
        <div className={classes.featureGrid}>
          <div className={classes.featureCard}>
            <div className={classes.featureIcon}>🎉</div>
            <h3>Discover Events</h3>
            <p>Find the perfect events happening in your area</p>
          </div>
          <div className={classes.featureCard}>
            <div className={classes.featureIcon}>👥</div>
            <h3>Connect</h3>
            <p>Meet new people who share your interests</p>
          </div>
          <div className={classes.featureCard}>
            <div className={classes.featureIcon}>📅</div>
            <h3>Organize</h3>
            <p>Create and manage your own events effortlessly</p>
          </div>
          <div className={classes.featureCard}>
            <div className={classes.featureIcon}>📍</div>
            <h3>Local Focus</h3>
            <p>Find events right in your neighborhood</p>
          </div>
        </div>
      </section>

      <section className={classes.localParties}>
        <h2>Happening Near You</h2>
        {loading ? (
          <div className={classes.loading}>Loading local events...</div>
        ) : localParties.length > 0 ? (
          <div className={classes.partyGrid}>
            {localParties.map((party) => (
              <div 
                key={party.id} 
                className={classes.partyCard}
                onClick={() => navigate(`/party/${party.id}`)}
              >
                {party.imageUrl && (
                  <div 
                    className={classes.partyImage}
                    style={{ backgroundImage: `url(${party.imageUrl})` }}
                  />
                )}
                <div className={classes.partyContent}>
                  <h3>{party.title}</h3>
                  <div className={classes.partyDetails}>
                    <span>📅 {new Date(party.date).toLocaleDateString()}</span>
                    <span>📍 {party.location}</span>
                    <span>👥 {party.attendees} attending</span>
                  </div>
                  <p>{party.description}</p>
                </div>
              </div>
            ))}
          </div>
        ) : (
          <div className={classes.noParties}>
            <p>No local events found at the moment.</p>
            <button 
              className={classes.primaryButton}
              onClick={() => navigate('/create-party')}
            >
              Create the First One
            </button>
          </div>
        )}
      </section>
    </div>
  );
};
