import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import classes from './HomePage.module.scss';
import { PartyPopulated } from "../../data/types/Party.ts";
import { useApi } from "../../context/ApiContext.ts";
import { toast } from "sonner";
import {convertPartiesToTableDatasource} from "../../data/utils/TableUtils.ts";
import {partyTableColumns, PartyTableRow} from "../../data/constants/TableColumns.tsx";
import {ActionButton, SortableTable} from "../../components/table/SortableTable.tsx";
import {useAppSelector} from "../../store/store-helper.ts";
import {isUserLoggedIn} from "../../store/slices/UserSlice.ts";
import AccessCodeModal from "./AccessCodeModal.tsx";
import {NavigateToPartyHome} from "../../data/utils/PartyUtils.ts";


export const Homepage = () => {
  const api = useApi();
  const navigate = useNavigate();
  const userLoggedIn = useAppSelector(isUserLoggedIn);
  const [publicParties, setPublicParties] = useState<PartyPopulated[]>([]);
  const [isAccessCodeModalVisible, setIsAccessCodeModalVisible] = useState(false);


    useEffect(() => {
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

  }, [api.partyApi]);

  const partyVisitLoggedOutActionButton: ActionButton<PartyTableRow> = {
    label: 'Visit',
    color: 'info',
    onClick: (party: PartyTableRow) => {
        NavigateToPartyHome(navigate, party.name, party.id, party.organizerName)
    }
  }

  const partyVisitLoggedInActionButton: ActionButton<PartyTableRow> = {
    label: 'Join',
    color: 'info',
    onClick: (party: PartyTableRow) => {
      api.partyAttendanceApi.joinPublicParty(party.id)
        .then(response => {
            if(response === 'error'){
              toast.error('Unable to join party')
              return
            }

            NavigateToPartyHome(navigate, party.name, party.id, party.organizerName)
        })
        .catch(() => {
            toast.error('Unexpected error')
        })
    }
  }


  return (
    <div className={classes.homepage}>
      <section className={classes.hero}>
        <div className={classes.heroContent}>
          <h1>Welcome to Party Organizer</h1>
          <p>Discover, create, and join amazing events in your area</p>
          {!userLoggedIn && (
              <div className={classes.heroButtons}>
                <button
                    className={classes.primaryButton}
                    onClick={() => navigate('/register')}
                >
                  Get started
                </button>
                <button
                    className={classes.secondaryButton}
                    onClick={() => navigate('/login')}
                >
                  Sign in
                </button>
              </div>)}
        </div>
      </section>

      <section className={classes.partiesSection}>
        <div className={classes.sectionContainer}>
          <h2>Public Parties</h2>
          <div className={classes.tableContainer}>
            <SortableTable
                columns={partyTableColumns}
                data={convertPartiesToTableDatasource(publicParties)}
                rowsPerPageOptions={[3, 5, 10, 15]}
                defaultRowsPerPage={5}
                  actionButtons={userLoggedIn ? [partyVisitLoggedInActionButton] : [partyVisitLoggedOutActionButton]}
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
            onClick={() => setIsAccessCodeModalVisible(true)}
          >
            <div className={classes.buttonContent}>
              <h3>Join Party</h3>
              <p>Find events to attend</p>
            </div>
          </button>
        </div>
      </section>
        <AccessCodeModal
            visible={isAccessCodeModalVisible}
            onClose={() => setIsAccessCodeModalVisible(false)}
        />
    </div>
  );
};
