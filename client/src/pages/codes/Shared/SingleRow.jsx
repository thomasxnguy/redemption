import React from "react";
import { Badge, Button } from "reactstrap";
import { useToasts } from "react-toast-notifications";

// TODO(Dan): Refactor to use Typescript
// interface Asset {
//   amount: Number;
//   token_id: String;
// }

// interface SingleCode {
//   link: String;
//   code: String;
//   provider: String;
//   valid: Boolean;
//   expiration_date: String;
//   created_date: String;
//   asset: {
//     coin: Number,
//     used: Boolean,
//     assets: [Asset]
//   };
// }

// interface AllCodesSProps {
//   code: SingleCode;
// }

const CopyLinkButton = ({ link }) => {
  const { addToast } = useToasts();
  const copyToClipboard = event => {
    navigator.clipboard.writeText(link).then(
      function() {
        addToast("Link copied to clipboard!", {
          appearance: "info",
          autoDismiss: true
        });
      },
      function(err) {
        addToast("Link copied to clipboard!", {
          appearance: "error",
          autoDismiss: true
        });
      }
    );
  };
  return (
    <div>
      <Button color="primary" size="sm" onClick={copyToClipboard}>
        Copy
      </Button>
    </div>
  );
};

export const SingleRow = props => {
  const { link, code, valid, asset } = props.code;
  const { used } = asset;
  return (
    <tr>
      <td>
        {valid ? (
          <Badge color="success">Valid</Badge>
        ) : (
          <Badge color="secondary">Invalid</Badge>
        )}
      </td>
      <td>
        {used ? (
          <Badge color="secondary">Redeemed</Badge>
        ) : (
          <Badge color="success">Active</Badge>
        )}
      </td>
      <td>
        <small>{code}</small>
      </td>
      <td>
        {asset.assets.map(({ amount, token_id }) => {
          return (
            <Button
              color="warning"
              className="btn-brand mb-1 mr-1"
              size="sm"
              key={token_id}
            >
              <i>{token_id}</i>
              <span>{amount} </span>
            </Button>
          );
        })}
      </td>
      <td>
        <CopyLinkButton link={link} />
      </td>
      <td>
        <Button color="success" size="sm">
          Details
        </Button>
      </td>
    </tr>
  );
};
