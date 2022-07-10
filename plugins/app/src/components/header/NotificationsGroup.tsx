import React, { useContext } from 'react';

import { IPluginsContext, PluginsContext } from '../../context/PluginsContext';
import { IGroup } from '../../context/NotificationsContext';
import { ITimes } from '@kobsio/shared';
import Module from '../module/Module';

interface INotificationsGroupProps {
  times: ITimes;
  group: IGroup;
}

const NotificationsGroup: React.FunctionComponent<INotificationsGroupProps> = ({
  group,
  times,
}: INotificationsGroupProps) => {
  const pluginsContext = useContext<IPluginsContext>(PluginsContext);
  const instance = pluginsContext.getInstance(group.plugin.satellite, group.plugin.type, group.plugin.name);

  const loadingContent = (): React.ReactElement => {
    return <div></div>;
  };

  const errorContent = (props: { title: string; children: React.ReactElement }): React.ReactElement => {
    return <div></div>;
  };

  if (!instance) {
    return <div></div>;
  }

  return (
    <Module
      version={pluginsContext.version}
      name={group.plugin.type}
      module="./Notifications"
      props={{
        instance: instance,
        options: group.plugin.options,
        times: times,
        title: group.title,
      }}
      errorContent={errorContent}
      loadingContent={loadingContent}
    />
  );
};

export default NotificationsGroup;
