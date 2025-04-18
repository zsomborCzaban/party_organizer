import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Table } from 'antd';
import classes from './HomePage.module.scss';
import { PartyPopulated } from "../data/types/Party.ts";
import { useApi } from "../context/ApiContext.ts";
import { toast } from "sonner";
import { partyTableColumnsLegacy } from '../data/constants/TableColumns';

export const Homepage = () => {
  const api = useApi();
  const navigate = useNavigate();
  const [publicParties, setPublicParties] = useState<PartyPopulated[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    setLoading(true);
    api.partyApi.getPublicParties()
      .then(resp => {
        if (resp === 'error') {
          toast.error('Unable to load public parties');
          return;
        }
        setPublicParties(resp.data);
      })
      .catch(() => {
        toast.error('Unexpected error');
      })
      .finally(() => {
        setLoading(false);
      });
  }, [api.partyApi]);

  const handlePartyClick = (party: PartyPopulated) => {
    navigate(`/party/${party.ID}`);
  };

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

      <section className={classes.partiesSection}>
        <div className={classes.sectionContainer}>
          <h2>Public Parties</h2>
          <div className={classes.tableContainer}>
            <Table
              dataSource={publicParties}
              columns={partyTableColumnsLegacy}
              loading={loading}
              onRow={(record) => ({
                onClick: () => handlePartyClick(record),
                style: { cursor: 'pointer' }
              })}
              pagination={{ pageSize: 5 }}
            />
          </div>
        </div>
      </section>

      <section className={classes.actionSection}>
        <div className={classes.actionButtons}>
          <button 
            className={classes.createButton}
            onClick={() => navigate('/createParty')}
          >
            <div className={classes.buttonContent}>
              <h3>Create Party</h3>
              <p>Start your own event</p>
            </div>
          </button>
          <button 
            className={classes.joinButton}
            onClick={() => navigate('/join-party')}
          >
            <div className={classes.buttonContent}>
              <h3>Join Party</h3>
              <p>Find events to attend</p>
            </div>
          </button>
        </div>
      </section>
    </div>
  );
};
