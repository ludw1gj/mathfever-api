import { GET_CATEGORIES, GET_CATEGORIES_ERROR } from './types';
import { getCategories as queryCategories } from '../../util/query';

/** Get categories. */
export const getCategories = () => {
  return (dispatch, _) => {
    queryCategories()
      .then(resp => {
        dispatch({ type: GET_CATEGORIES, payload: resp.data });
      })
      .catch(err => {
        dispatch({ type: GET_CATEGORIES_ERROR, payload: err });
      });
  };
};
