import { forwardRef, type InputHTMLAttributes } from 'react';

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  /** Error state */
  error?: boolean;
}

export const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ className = '', error = false, ...props }, ref) => {
    return (
      <input
        ref={ref}
        className={`w-full rounded-xl border bg-white px-4 py-3 text-sm text-slate-900 placeholder:text-slate-400 transition-colors focus:outline-none focus:ring-2 focus:ring-offset-0 disabled:cursor-not-allowed disabled:opacity-50 dark:bg-zinc-900 dark:text-white dark:placeholder:text-zinc-500 ${
          error
            ? 'border-red-300 focus:border-red-500 focus:ring-red-500/20 dark:border-red-800'
            : 'border-slate-200 focus:border-[#1e3a8a] focus:ring-[#1e3a8a]/20 dark:border-zinc-700 dark:focus:border-blue-500'
        } ${className}`}
        {...props}
      />
    );
  }
);

Input.displayName = 'Input';
