import { combineReducers } from 'redux';

import { categoriesReducer } from './categoriesReducer';

/** The root reducer. */
const rootReducer = combineReducers({
  categories: categoriesReducer
});

export default rootReducer;
