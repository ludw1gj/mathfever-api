import { GET_CATEGORIES, GET_CATEGORIES_ERROR } from '../actions/types';

const initState = {
  data: null,
  error: null
};

/** Reducer containing categories data. */
export const categoriesReducer = (state = initState, action) => {
  switch (action.type) {
    case GET_CATEGORIES:
      return {
        data: action.payload,
        error: null
      };

    case GET_CATEGORIES_ERROR:
      return {
        ...state,
        error: action.payload
      };

    default:
      return state;
  }
};
