import { type ReactNode } from 'react';

interface FormFieldProps {
  /** Field label */
  label?: string;
  /** Error message */
  error?: string;
  /** Helper text */
  hint?: string;
  /** Required indicator */
  required?: boolean;
  /** HTML for attribute */
  htmlFor?: string;
  /** Form field input */
  children: ReactNode;
  /** Additional classes */
  className?: string;
}

export function FormField({
  label,
  error,
  hint,
  required = false,
  htmlFor,
  children,
  className = '',
}: FormFieldProps) {
  return (
    <div className={`space-y-1.5 ${className}`}>
      {label && (
        <label
          htmlFor={htmlFor}
          className="block text-sm font-medium text-slate-700 dark:text-zinc-300"
        >
          {label}
          {required && <span className="ml-1 text-[#dc2626]">*</span>}
        </label>
      )}
      {children}
      {error && (
        <p className="text-xs text-[#dc2626] dark:text-red-400">{error}</p>
      )}
      {hint && !error && (
        <p className="text-xs text-slate-500 dark:text-zinc-500">{hint}</p>
      )}
    </div>
  );
}
