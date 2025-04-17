import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../../store/store.ts';
import { closeDrawer } from '../../store/slices/profileDrawerSlice.ts';
import {
  Drawer,
  DrawerContent,
  DrawerHeader,
  DrawerTitle,
  DrawerDescription,
  DrawerClose,
} from '../ui/drawer.tsx';
import { Button } from '../ui/button.tsx';

export const ProfileDrawer = () => {
  const dispatch = useDispatch();
  const isOpen = useSelector((state: RootState) => state.profileDrawer.isOpen);

  return (
    <Drawer open={isOpen} onOpenChange={(open) => !open && dispatch(closeDrawer())}>
      <DrawerContent>
        <DrawerHeader>
          <DrawerTitle>Profile</DrawerTitle>
          <DrawerDescription>Your profile information</DrawerDescription>
        </DrawerHeader>
        <div className="p-4">
          <p>Profile content goes here</p>
        </div>
        <div className="p-4">
          <DrawerClose asChild>
            <Button variant="outline">Close</Button>
          </DrawerClose>
        </div>
      </DrawerContent>
    </Drawer>
  );
}; 