import { ArrowDownward, ArrowUpward, ContentCopy, Delete } from '@mui/icons-material';
import {
  Box,
  Divider,
  IconButton,
  List,
  ListItem,
  ListItemButton,
  ListItemText,
  Stack,
  TextField,
  Typography,
} from '@mui/material';
import { FunctionComponent, useState } from 'react';

interface ILogsFieldsList {
  fields: string[];
  onSwapItem: (from: number, to: number) => void;
  onToggleField: (key: string) => void;
  selectedFields: string[];
}

/**
 * LogsFieldsList renders the given fields ina list format
 * fields are toggle'able and trigger the callback `onToggleField`
 */
const LogsFieldsList: FunctionComponent<ILogsFieldsList> = ({ fields, selectedFields, onSwapItem, onToggleField }) => {
  const [search, setSearch] = useState('');
  return (
    <>
      <List
        sx={{
          position: 'relative',
          wordBreak: 'break-all',
        }}
        dense={true}
        subheader={<li />}
      >
        {selectedFields.length > 0 && (
          <>
            <Typography variant="h5" p={2}>
              Selected fields
            </Typography>
            <Divider />
          </>
        )}
        {selectedFields
          .filter((field) => field !== 'timestamp')
          .map((field, i) => (
            <ListItem
              key={field}
              disablePadding={true}
              aria-label={field}
              sx={{
                '&:hover .fields-action-items': { opacity: 1 },
              }}
            >
              <ListItemText primary={field} sx={{ mx: 4 }} />
              <Stack
                className="fields-action-items"
                direction="row"
                sx={{
                  mr: 2,
                  opacity: 0,
                }}
              >
                <IconButton
                  aria-label="move down"
                  sx={{ p: 0 }}
                  size="small"
                  disabled={i === selectedFields.length - 1}
                  disableRipple={true}
                  onClick={() => {
                    onSwapItem(i, i + 1);
                  }}
                >
                  <ArrowDownward sx={{ fontSize: 14 }} />
                </IconButton>
                <IconButton
                  aria-label="move up"
                  sx={{ p: 0 }}
                  size="small"
                  disabled={i === 0}
                  disableRipple={true}
                  onClick={() => {
                    onSwapItem(i - 1, i);
                  }}
                >
                  <ArrowUpward sx={{ fontSize: 14 }} />
                </IconButton>
                <IconButton
                  aria-label="copy"
                  sx={{ p: 0 }}
                  size="small"
                  disableRipple={true}
                  onClick={() => {
                    navigator.clipboard.writeText(field);
                  }}
                >
                  <ContentCopy sx={{ fontSize: 14 }} />
                </IconButton>
                <IconButton
                  aria-label="delete"
                  disableRipple={true}
                  onClick={() => onToggleField(field)}
                  sx={{ p: 0 }}
                  size="small"
                >
                  <Delete sx={{ fontSize: 14 }} />
                </IconButton>
              </Stack>
            </ListItem>
          ))}
        <Typography variant="h5" p={2}>
          Fields
        </Typography>
        <Divider />
        <Box p={2}>
          <TextField label="search field" size="small" onChange={(e) => setSearch(e.target.value)} fullWidth={true} />
        </Box>
        {fields
          .filter((field) => field !== 'timestamp')
          .filter((field) => field.includes(search))
          .map((field) => (
            <ListItem key={field} disablePadding={true}>
              <ListItemButton onClick={() => onToggleField(field)} aria-label={field}>
                <ListItemText primary={field} />
              </ListItemButton>
            </ListItem>
          ))}
      </List>
    </>
  );
};

export default LogsFieldsList;