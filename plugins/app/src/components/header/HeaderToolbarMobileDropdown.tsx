import { Dropdown, DropdownItem, DropdownSeparator, KebabToggle } from '@patternfly/react-core';
import React, { useContext } from 'react';
import CogIcon from '@patternfly/react-icons/dist/esm/icons/cog-icon';
import { Link } from 'react-router-dom';
import { QuestionCircleIcon } from '@patternfly/react-icons/dist/esm/icons/question-circle-icon';

import { AuthContext, IAuthContext } from '../../context/AuthContext';

interface IHeaderToolbarMobileDropdownProps {
  isOpen: boolean;
  setIsOpen: (value: boolean) => void;
}

const HeaderToolbarMobileDropdown: React.FunctionComponent<IHeaderToolbarMobileDropdownProps> = ({
  isOpen,
  setIsOpen,
}: IHeaderToolbarMobileDropdownProps) => {
  const authContext = useContext<IAuthContext>(AuthContext);

  if (authContext.user.email) {
    return (
      <Dropdown
        isPlain={true}
        position="right"
        toggle={<KebabToggle onToggle={(): void => setIsOpen(!isOpen)} />}
        isOpen={isOpen}
        dropdownItems={[
          // eslint-disable-next-line @typescript-eslint/ban-ts-comment
          // @ts-ignore
          <DropdownItem key="myprofile" component={(props): React.ReactNode => <Link {...props} to="/profile" />}>
            My profile
          </DropdownItem>,
          <DropdownSeparator key="divider1" />,
          <DropdownItem
            key="logout"
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore
            // eslint-disable-next-line jsx-a11y/anchor-has-content
            component={(props): React.ReactElement => <a {...props} href="/api/auth/logout" />}
          >
            Logout
          </DropdownItem>,
          <DropdownSeparator key="divider2" />,
          // eslint-disable-next-line @typescript-eslint/ban-ts-comment
          // @ts-ignore
          <DropdownItem key="settings" component={(props): React.ReactElement => <Link {...props} to="/settings" />}>
            <CogIcon /> Settings
          </DropdownItem>,
          <DropdownItem
            key="help"
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore
            component={(props): React.ReactElement => (
              <a {...props} href="https://kobs.io" target="_blank" rel="noreferrer">
                <QuestionCircleIcon /> Help
              </a>
            )}
          ></DropdownItem>,
        ]}
      />
    );
  }

  return (
    <Dropdown
      isPlain={true}
      position="right"
      toggle={<KebabToggle onToggle={(): void => setIsOpen(!isOpen)} />}
      isOpen={isOpen}
      dropdownItems={[
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        <DropdownItem key="settings" component={(props): React.ReactElement => <Link {...props} to="/settings" />}>
          <CogIcon /> Settings
        </DropdownItem>,
        <DropdownItem
          key="help"
          // eslint-disable-next-line @typescript-eslint/ban-ts-comment
          // @ts-ignore
          component={(props): React.ReactElement => (
            <a {...props} href="https://kobs.io" target="_blank" rel="noreferrer">
              <QuestionCircleIcon /> Help
            </a>
          )}
        ></DropdownItem>,
      ]}
    />
  );
};

export default HeaderToolbarMobileDropdown;
