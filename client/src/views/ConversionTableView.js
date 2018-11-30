import React from 'react';

import DocumentTitle from '../components/DocumentTitle';
import PageHeading from '../components/PageHeading';
import ContentCard from '../components/ContentCard';
import GenericDocMetaDesc from '../components/GenericDocMetaDesc';

/** Conversion Table page. */
const ConversionTableView = () => {
  return (
    <div>
      <DocumentTitle>Conversion Table - MathFever</DocumentTitle>
      <GenericDocMetaDesc />

      <PageHeading>Networking</PageHeading>

      <ContentCard>
        <h4>Conversion Table</h4>

        <table className="mdl-data-table mdl-js-data-table">
          <tbody>
            <tr>
              <th className="mdl-data-table__cell--non-numeric">Binary</th>
              <th className="mdl-data-table__cell--non-numeric">Decimal</th>
              <th className="mdl-data-table__cell--non-numeric">Hexadecimal</th>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (0000)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (0)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (0)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (0001)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (1)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (1)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (0010)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (2)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (2)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (0011)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (3)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (3)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (0100)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (4)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (4)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (0101)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (5)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (5)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (0110)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (6)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (6)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (0111)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (7)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (7)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (1000)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (8)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (8)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (1001)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (9)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (9)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (1010)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (10)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (A)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (1011)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (11)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (B)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (1100)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (12)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (C)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (1101)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (13)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (D)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (1110)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (14)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (E)
                <sub>16</sub>
              </td>
            </tr>
            <tr>
              <td className="mdl-data-table__cell--non-numeric">
                (1111)
                <sub>2</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (15)
                <sub>10</sub>
              </td>
              <td className="mdl-data-table__cell--non-numeric">
                (F)
                <sub>16</sub>
              </td>
            </tr>
          </tbody>
        </table>
      </ContentCard>
    </div>
  );
};

export default ConversionTableView;
